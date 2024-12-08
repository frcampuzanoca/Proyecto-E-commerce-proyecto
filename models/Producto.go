package models

import (
	"errors"
)

// Producto representa un artículo disponible para la venta
type Producto struct {
	ID          int      `json:"id"`
	Nombre      string   `json:"nombre"`
	Descripcion string   `json:"descripcion"`
	Precio      float64  `json:"precio"`
	Categorias  []string `json:"categorias"`
	Stock       int      `json:"stock"`
	Imagen      string   `json:"imagen"`
}

// Devuelve y asigna el precio del producto
func (p *Producto) GetPrecio() float64 {
	return p.Precio
}

func (p *Producto) SetPrecio(precio float64) {
	p.Precio = precio
}

// GetStock y SetStock devuelve y asigna el stock del producto
func (p *Producto) GetStock() int {
	return p.Stock
}

func (p *Producto) SetStock(stock int) {
	p.Stock = stock
}

// Get y SetDescripcion actualiza y asigna la descripción del producto
func (p *Producto) GetDescripcion() string {
	return p.Descripcion
}

func (p *Producto) SetDescripcion(descripcion string) {
	p.Descripcion = descripcion
}

// Validación y asignación del precio
func (p *Producto) ValidarPrecio(precio float64) error {
	if precio <= 0 {
		return errors.New("El precio no puede ser igual o menor a cero.")
	}
	p.Precio = precio
	return nil
}

// ActualizarStock actualiza el inventario del producto
func (p *Producto) ActualizarStock(cantidad int) error {
	if p.Stock+cantidad < 0 {
		return errors.New("Los productos en stock no pueden ser menos de cero")
	}
	p.Stock += cantidad
	return nil
}
