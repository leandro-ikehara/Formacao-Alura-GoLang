package main

import (
	"pizzaria/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//nomePizzaria := "Go Pizza"
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.Run()
}

func getPizzas(c *gin.Context) {
	var pizzas = []models.Pizza{
		{ID: 1, Nome: "Margherita", Preco: 25.50},
		{ID: 2, Nome: "Pepperoni", Preco: 30.00},
		{ID: 3, Nome: "Quattro Formaggi", Preco: 35.00},
	}
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}
