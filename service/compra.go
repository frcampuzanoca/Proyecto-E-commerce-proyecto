package service

import (
	"Proyecto-E-commerce-proyecto/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func CrearCompraHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		CrearCompra(c, db)
	}
}

func CrearCompra(c *gin.Context, db *sql.DB) {

	var nuevaCompra models.Compra

	// Validar el JSON recibido
	if err := c.ShouldBindJSON(&nuevaCompra); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Guardar la compra (simulado, puedes usar una base de datos aquí)
	nuevaCompra.Estado = "completado" // Marcar la compra como completada

	// Simular guardar la compra en una base de datos
	// Aquí puedes insertar los datos en una base de datos relacional o no relacional
	// Por ejemplo: MySQL, PostgreSQL, MongoDB, etc.
	guardarCompra(db, nuevaCompra)

	c.JSON(http.StatusOK, gin.H{
		"message": "Compra realizada con éxito",
		"compra":  nuevaCompra,
	})

}

func guardarCompra(db *sql.DB, compra models.Compra) error {
	// Iniciar una transacción
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Insertar la compra
	queryCompra := `INSERT INTO compras (nombre, correo, direccion, total, estado, fecha, cantidad_items) VALUES (?, ?, ?, ?, ?, ?, ?)`
	res, err := tx.Exec(queryCompra, compra.Nombre, compra.Correo, compra.Direccion, compra.Total, compra.Estado, time.Now(), len(compra.Productos))
	if err != nil {
		tx.Rollback()
		return err
	}

	// Obtener el ID de la compra recién insertada
	compraID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insertar los productos de la compra
	queryProducto := `INSERT INTO productos_comprados (compra_id, nombre, precio, cantidad, subtotal) VALUES (?, ?, ?, ?, ?)`
	for _, producto := range compra.Productos {
		_, err := tx.Exec(queryProducto, compraID, producto.Nombre, producto.Precio, producto.Cantidad, producto.Subtotal)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Confirmar la transacción
	err = tx.Commit()
	if err != nil {
		return err
	}

	log.Printf("Compra guardada con éxito: %+v", compra)
	return nil
}

func ObtenerCompras(db *sql.DB) ([]models.Compra, error) {
	query := `
        SELECT 
            c.fecha, c.nombre, c.correo, c.total, c.estado, 
            COALESCE(SUM(p.cantidad), 0) AS cantidad_items
        FROM 
            compras c
        LEFT JOIN 
            productos_comprados p ON c.id = p.compra_id
        GROUP BY 
            c.id
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var compras []models.Compra
	for rows.Next() {
		var compra models.Compra
		err := rows.Scan(&compra.Fecha, &compra.Nombre, &compra.Correo, &compra.Total, &compra.Estado, &compra.CantidadItems)
		if err != nil {
			return nil, err
		}
		compras = append(compras, compra)
	}

	return compras, nil
}
