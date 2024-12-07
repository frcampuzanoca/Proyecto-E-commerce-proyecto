package models

import "errors"

// Pedido representa una orden de compra
type Pedido struct {
	ID           int           `json:"id"`
	ClienteID    int           `json:"cliente_id"`
	Productos    []Producto    `json:"productos"`
	Total        float64       `json:"total"`
	Estado       string        `json:"estado"`
	Pago         Pago          `json:"pago"`
	Calificacion *Calificacion `json:"calificacion,omitempty"`
}

// GetEstado y SetEstado devuelve y actualzia  el estado del pedido
func (p *Pedido) GetEstado() string {
	return p.Estado
}

func (p *Pedido) SetEstado(estado string) {
	p.Estado = estado
}

// GetProductos devuelve la lista de productos del pedido
func (p *Pedido) GetProductos() []Producto {
	return p.Productos
}

// SetPago asigna un pago al pedido
func (p *Pedido) SetPago(pago Pago) {
	p.Pago = pago
}

// SetTotal actualiza el total del pedido
func (p *Pedido) SetTotal(total float64) error {
	if total < 0 {
		return errors.New("El valor total no puede ser menor a cero")
	}
	p.Total = total
	return nil
}

// ActualizarEstado cambia el estado del pedido
func (p *Pedido) ActualizarEstado(nuevoEstado string) {
	p.Estado = nuevoEstado
}

func (p *Pedido) EstadoPedido(estado string) error {
	validarEstado := []string{"Completado", "Cancelado", "En proceso", "Pendiente"}
	for _, s := range validarEstado {
		if estado == s {
			p.Estado = estado
			return nil
		}
	}
	return errors.New("Estado incorrecto")
}
