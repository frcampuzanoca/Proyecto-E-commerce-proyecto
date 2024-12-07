package models

import "time"

// Calificacion representa una evaluación de un pedido
type Calificacion struct {
	ID         int       `json:"id"`
	PedidoID   int       `json:"pedido_id"`
	Puntuacion int       `json:"puntuacion"`
	Comentario string    `json:"comentario"`
	Fecha      time.Time `json:"fecha"`
}

// GetPuntuacion  y Set devuelve y actualiza la puntuación de la calificación
func (c *Calificacion) GetPuntuacion() int {
	return c.Puntuacion
}

func (c *Calificacion) SetPuntuacion(puntuacion int) {
	c.Puntuacion = puntuacion
}

// DarCalificacion asigna puntuación y comentario a una calificación
func (c *Calificacion) DarCalificacion(puntuacion int, comentario string) {
	c.Puntuacion = puntuacion
	c.Comentario = comentario
	c.Fecha = time.Now()
}
