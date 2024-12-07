package models

import (
	"errors"
	"time"
)

// Pago contiene los detalles de una transacción
type Pago struct {
	ID     int       `json:"id"`
	Monto  float64   `json:"monto"`
	Fecha  time.Time `json:"fecha"`
	Metodo string    `json:"metodo"`
	Estado string    `json:"estado"`
}

// Devuelve el Id del pago
func (p *Pago) GetID() int {
	return p.ID
}

// GetEstado y setEstado devuelve y asigna el estado del pago
func (p *Pago) GetEstado() string {
	return p.Estado
}

func (p *Pago) SetEstado(estado string) {
	p.Estado = estado
}

// Get Metodo y SetMetodo asigna el método de pago
func (p *Pago) GetMetodo() string {
	return p.Metodo
}

func (p *Pago) SetMetodo(metodo string) {
	p.Metodo = metodo
}

// Validación y asignación para que el monto del pago no sea menor o igual a cero.
func (p *Pago) SetMonto(monto float64) error {
	if monto <= 0 {
		return errors.New("El monto no puede ser cero o menor a cero.")
	}
	p.Monto = monto
	return nil
}

// RealizarPago simula un pago exitoso
func (p *Pago) RealizarPago() error {
	if p.Estado == "Completado" {
		return errors.New("El pago ya fue realizado")
	}
	p.Estado = "Completado"
	p.Fecha = time.Now()
	return nil
}
