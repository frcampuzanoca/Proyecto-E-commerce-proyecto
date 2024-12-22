package service

import (
	"Proyecto-E-commerce-proyecto/models"
	"database/sql"
	"log"
)

// GetAllProducts obtiene todos los productos de la base de datos
func GetAllProducts(db *sql.DB) ([]models.Producto, error) {
	query := `
        SELECT id, nombre, categoria, precio, descripcion, stock, fecha_creacion 
        FROM productos
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []models.Producto
	for rows.Next() {
		var p models.Producto
		err := rows.Scan(
			&p.ID,
			&p.Nombre,
			&p.Categoria,
			&p.Precio,
			&p.Descripcion,
			&p.Stock,
			&p.Imagen,
			&p.FechaCreacion,
		)
		if err != nil {
			log.Printf("Error al escanear fila: %v", err)
			continue // Puedes decidir si continuar o retornar el error
		}
		productos = append(productos, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productos, nil
}
