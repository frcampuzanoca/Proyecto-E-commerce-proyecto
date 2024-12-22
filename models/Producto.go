// models/producto.go
package models

import "time"

// Producto representa un producto en la tienda
type Producto struct {
	ID            int       `json:"id"`
	Nombre        string    `json:"nombre"`
	Categoria     string    `json:"categoria"`
	Precio        float64   `json:"precio"`
	Descripcion   string    `json:"descripcion,omitempty"`
	Stock         int       `json:"stock"`
	Imagen        string    `json:"imagen"`
	FechaCreacion time.Time `json:"fecha_creacion"`
}
