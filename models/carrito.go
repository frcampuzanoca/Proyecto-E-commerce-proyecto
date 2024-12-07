package models

import (
	"errors"
)

// Carrito almacena productos añadidos por un cliente
type Carrito struct {
	ID        int        `json:"id"`
	Productos []Producto `json:"productos"`
	Total     float64    `json:"total"`
}

// GetID devuelve el ID del carrito
func (c *Carrito) GetID() int {
	return c.ID
}

// GetProductos y SetProductos Devuelv y asigna los productos del carrito
func (c *Carrito) GetProductos() []Producto {
	return c.Productos
}

func (c *Carrito) SetProductos(productos []Producto) {
	c.Productos = productos
}

// Actualiza el total del carrito
func (c *Carrito) SetTotal(total float64) error {
	if total < 0 {
		return errors.New("El total no puede ser menor a cero")
	}
	c.Total = total
	return nil
}

// AgregarProducto añade un producto al carrito y recalcula el total
func (c *Carrito) AgregarProducto(producto Producto) {
	c.Productos = append(c.Productos, producto)
	c.CalcularTotal()
}

// EliminarProducto elimina un producto del carrito por su ID
func (c *Carrito) EliminarProducto(idProducto int) {
	for i, p := range c.Productos {
		if p.ID == idProducto {
			c.Productos = append(c.Productos[:i], c.Productos[i+1:]...)
			break
		}
	}
	c.CalcularTotal()
}

// CalcularTotal calcula el costo total de los productos en el carrito
func (c *Carrito) CalcularTotal() {
	var total float64
	for _, p := range c.Productos {
		total += p.Precio
	}
	c.Total = total
}
