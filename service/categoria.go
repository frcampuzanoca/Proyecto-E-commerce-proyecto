package service

import (
	"Proyecto-E-commerce-proyecto/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FormularioCrearCategoria(c *gin.Context) {
	c.HTML(http.StatusOK, "crear-categoria.html", nil)
}

func CrearCategoria(c *gin.Context) {

	var nuevaCategoria models.Categoria

	// Obtener datos del formulario
	nuevaCategoria.Nombre = c.PostForm("nombre")
	nuevaCategoria.ID = 1

	c.HTML(http.StatusOK, "ver-categoria.html", gin.H{
		"Nombre": nuevaCategoria.Nombre,
		"ID":     nuevaCategoria.ID,
	})
}
