package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func main() {
	loadPizzas()
	//nomePizzaria := "Go Pizza"
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizzasByID)
	router.Run()
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func getPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID must be a number"})
		return
	}
	for _, pizza := range pizzas {
		if pizza.ID == id {
			c.JSON(200, pizza)
			return
		}
	}
	c.JSON(404, gin.H{"error": "Pizza not found"})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	savePizza()
	c.JSON(201, newPizza)
}

func loadPizzas() {
	file, err := os.Open("dados/pizza.json")
	if err != nil {
		fmt.Println("Error file: ", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error decoding JSON: ", err)
		return
	}
}

func savePizza() {
	file, err := os.Create("dados/pizza.json")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error encoding JSON: ", err)
		return
	}
}
