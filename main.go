package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// car represents data about each car's record.
type car struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

var cars = []car{
	{ID: "1", Title: "BMW", Color: "Black"},
	{ID: "2", Title: "Tesla", Color: "Red"},
}

func main() {
	router := gin.Default()
	router.GET("/cars", getCars)
	router.GET("/car/:id", getCarByID)
	router.POST("/car", createCar)
	router.DELETE("/car/:id", deleteCar)
	router.PUT("/car/:id", replaceCar)

	router.Run("localhost:8080")
}

func createCar(context *gin.Context) {
	var newCar car
	if err := context.BindJSON(&newCar); err != nil {
		return
	}

	cars = append(cars, newCar)
	context.IndentedJSON(http.StatusCreated, newCar)
}

func getCars(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, cars)
}

func getCarByID(c *gin.Context) {
	for _, car := range cars {
		if car.ID == c.Param("id") {
			c.IndentedJSON(http.StatusOK, car)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": " Requested car is not found"})
}

func deleteCar(c *gin.Context) {
	for index, car := range cars {
		if car.ID == c.Param("id") {
			cars = append(cars[:index], cars[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": " Car is deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": " Car is not found"})
}

func replaceCar(c *gin.Context) {
	var newCar car
	if err := c.BindJSON(&newCar); err != nil {
		return
	}

	for index, car := range cars {
		if car.ID == c.Param("id") {
			cars = append(cars[:index], cars[index+1:]...)
			cars = append(cars, newCar)
			c.IndentedJSON(http.StatusOK, gin.H{"message": " Car is replaced"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": " Car is not found"})
}
