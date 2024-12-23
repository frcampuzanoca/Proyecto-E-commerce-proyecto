package models

type Compra struct {
	Nombre    string     `json:"nombre" binding:"required"`
	Correo    string     `json:"correo" binding:"required,email"`
	Direccion string     `json:"direccion" binding:"required"`
	Productos []Producto `json:"productos" binding:"required"`
	Total     float64    `json:"total" binding:"required"`
	Estado    string     `json:"estado"`
}

type ProductoCompra struct {
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
	Subtotal float64 `json:"subtotal"`
}
