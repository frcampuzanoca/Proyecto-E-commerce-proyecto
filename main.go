// Proyecto final
//Nombres: Cristhian Sare, Freddy Campuzano, Iriany Colina
/*Descripción: Avance pequeño donde podemos observar la estructura de los
productos y del usuario junto con la función getProduct*/
//Materia: Programación orientada a objetos.

// main.go
// main.go
package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	database "Proyecto-E-commerce-proyecto/Database"
	"Proyecto-E-commerce-proyecto/models"
	"Proyecto-E-commerce-proyecto/service"
)

// Productos de ejemplo y carrito global
var (
	productos = []models.Producto{
		{
			ID:     1,
			Nombre: "Laptop",
			Precio: 1000.0,
			Stock:  10,
			Imagen: "https://i.dell.com/is/image/DellContent/content/dam/ss2/product-images/page/category/laptop/xps/fy24-family-launch/prod-312204-laptop-xps-16-9640-14-9440-13-9340-sl-800x620.png?fmt=png-alpha&wid=800&hei=620",
		},
		{
			ID:     2,
			Nombre: "Smartphone",
			Precio: 700.0,
			Stock:  15,
			Imagen: "https://cdsassets.apple.com/live/7WUAS350/images/tech-specs/iphone-16.png",
		},
		{
			ID:     3,
			Nombre: "Headphones",
			Precio: 100.0,
			Stock:  20,
			Imagen: "https://i.ytimg.com/vi/UK72L99YbUE/hq720.jpg?sqp=-oaymwEhCK4FEIIDSFryq4qpAxMIARUAAAAAGAElAADIQj0AgKJD&rs=AOn4CLDmMCosMtxXMOxdtvoUClYWQUbUPQ",
		},
	}
	carrito = models.Carrito{ID: 1}
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error al cargar .env: %v", err)
	}

	// Conectar a la base de datos
	db, errDB := database.Connect()
	if errDB != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", errDB)
	}
	defer db.Close()

	// Inicializar Gin
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// Rutas
	r.GET("/", func(c *gin.Context) {
		// Página principal
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/ver-productos", func(c *gin.Context) {
		VerProductos(c, db)
	})

	// Iniciar el servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	//Página principal
	r.GET("/", func(c *gin.Context) {
		if db != nil {
			productosDB, err := service.GetAllProducts(db)
			if err == nil {
				productos = productosDB
			}
		}
		c.HTML(http.StatusOK, "index.html", gin.H{"productos": productos, "totalCarrito": carrito.Total})
	})

	// Rutas de carrito
	r.POST("/add-to-cart", func(c *gin.Context) {
		// Implementación básica como en el ejemplo anterior
	})
	r.GET("/cart", func(c *gin.Context) {
		// Implementación básica como en el ejemplo anterior
	})

	r.GET("/ver-productos", func(c *gin.Context) {
		VerProductos(c, db)
	})
	// Iniciar el servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

// VerProductos obtiene los productos desde la base de datos y los devuelve en JSON
func VerProductos(c *gin.Context, db *sql.DB) {
	productos, err := service.GetAllProducts(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener productos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"productos": productos})
}
