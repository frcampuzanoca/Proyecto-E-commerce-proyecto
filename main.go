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
	"Proyecto-E-commerce-proyecto/service"
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

		productos, err := service.GetAllProducts(db)
		if err != nil {
			c.JSON(500, "errror al obtener los productos")
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"productos": productos,
		})
	})

	r.GET("/cart", func(context *gin.Context) {
		context.HTML(http.StatusOK, "cart.html", gin.H{})
	})

	r.POST("/api/comprar", service.CrearCompraHandler(db))

	r.GET("/compras", func(context *gin.Context) {
		compras, err := service.ObtenerCompras(db)
		if err != nil {
			log.Printf("Error al obtener compras desde la DB: %v", err)
		}

		context.HTML(http.StatusOK, "compras.html", gin.H{
			"compras": compras,
		})
	})

	r.GET("/productos", func(context *gin.Context) {
		productos, err := service.GetAllProducts(db)

		if err != nil {
			context.JSON(500, "errror al obtener los productos")
		}
		context.HTML(http.StatusOK, "gestion_productos.html", gin.H{
			"productos": productos,
		})
	})

	// Iniciar el servidor
	r.Run(":8080")
}
