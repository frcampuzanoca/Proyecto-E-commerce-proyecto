// Avance del Autonomo#3
//Nombres: Cristhian Sare, Freddy Campuzano, Iriany
/*Descripción: Avance pequeño donde podemos observar la estructura de los
productos y del usuario junto con la función getProduct*/
//Materia: Programación orientada a objetos.

package main

import (
	"Proyecto-E-commerce-proyecto/models"
	"Proyecto-E-commerce-proyecto/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	productos = []models.Producto{
		{ID: 1, Nombre: "Laptop", Precio: 1000.0, Stock: 10},
		{ID: 2, Nombre: "Smartphone", Precio: 700.0, Stock: 15},
		{ID: 3, Nombre: "Headphones", Precio: 100.0, Stock: 20},
	}
	carrito = models.Carrito{ID: 1}
)

func main() {
	r := gin.Default()

	// Cargar plantillas HTML
	r.LoadHTMLGlob("templates/*")

	// Servir archivos estáticos (CSS, JS)
	r.Static("/static", "./static")

	r.GET("/formulario-categoria", service.FormularioCrearCategoria)
	r.POST("/crear-categoria", service.CrearCategoria)

	// Ruta para mostrar la página principal del e-commerce
	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"productos":    productos,
			"totalCarrito": carrito.Total,
		})
	})

	// Ruta para agregar un producto al carrito
	r.POST("/add-to-cart", service.AddToCart)

	// Ruta para mostrar el contenido del carrito
	r.GET("/cart", service.GetCart)

	// Iniciar el servidor
	r.Run(":8080")
}
