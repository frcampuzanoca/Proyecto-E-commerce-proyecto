package service

import (
	"database/sql"
	"fmt"

	"Proyecto-E-commerce-proyecto/models"
)

// GetAllProducts (exportada, empieza con mayúscula) hace SELECT a la tabla "productos".
func GetAllProducts(db *sql.DB) ([]models.Producto, error) {
	rows, err := db.Query(`
        SELECT 
            id,
            nombre,
            precio,
            stock,
            imagen
        FROM productos
    `)
	if err != nil {
		return nil, fmt.Errorf("error al consultar productos: %v", err)
	}
	defer rows.Close()

	var productos []models.Producto

	for rows.Next() {
		var p models.Producto
		if err := rows.Scan(&p.ID, &p.Nombre, &p.Precio, &p.Stock, &p.Imagen); err != nil {
			return nil, fmt.Errorf("error al escanear producto: %v", err)
		}
		productos = append(productos, p)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error en el recorrido de productos: %v", err)
	}

	return productos, nil
}
