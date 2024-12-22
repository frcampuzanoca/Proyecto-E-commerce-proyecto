// Avance del Autonomo#3
//Nombres: Cristhian Sare, Freddy Campuzano, Iriany
/*Descripción: Avance pequeño donde podemos observar la estructura de los
productos y del usuario junto con la función getProduct*/
//Materia: Programación orientada a objetos.

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	database "Proyecto-E-commerce-proyecto/Database"
	"Proyecto-E-commerce-proyecto/models"
	"Proyecto-E-commerce-proyecto/service"
)

// Productos “de presentación” (por si falla la BD o los quieres mostrar de ejemplo)
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
	// 2. Cargar las variables de entorno del archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error al cargar .env: %v", err)
		// No es forzosamente un "fatal"; puede seguir corriendo sin .env
	}

	// 3. (Opcional) Conectarte a la base de datos
	db, errDB := database.Connect()
	if errDB != nil {
		log.Printf("Error al conectar con la base de datos: %v", errDB)
	} else {
		defer db.Close()
	}

	// 4. Inicializar Gin
	r := gin.Default()

	// Cargar plantillas HTML
	r.LoadHTMLGlob("templates/*")

	// Servir archivos estáticos (CSS, JS)
	r.Static("/static", "./static")

	// Rutas para formulario y creación de categorías
	r.GET("/formulario-categoria", service.FormularioCrearCategoria)
	r.POST("/crear-categoria", service.CrearCategoria)

	// Ruta para la página principal del e-commerce
	r.GET("/", func(c *gin.Context) {
		// Aquí podrías, si quieres, intentar obtener productos de la BD:

		if db != nil {
			productosDB, err := service.GetAllProducts(db)
			if err != nil {
				log.Printf("Error al obtener productos desde la DB: %v", err)
			} else {
				productos = productosDB // Sobrescribes o combinas
			}
		}

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
