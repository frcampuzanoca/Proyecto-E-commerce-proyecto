package service

import (
	"Proyecto-E-commerce-proyecto/models"
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

func GetCart(c *gin.Context) {
	c.HTML(http.StatusOK, "cart.html", gin.H{
		"productos": carrito.Productos,
		"total":     carrito.Total,
	})

}

func AddToCart(c *gin.Context) {
	var productoID int
	if err := c.ShouldBindJSON(&productoID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto inválido"})
		return
	}

	// Buscar producto por ID
	for _, p := range productos {
		if p.ID == productoID {
			carrito.AgregarProducto(p)
			c.JSON(http.StatusOK, gin.H{"message": "Producto agregado al carrito"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
}
