package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/services", getServices)
	router.POST("/services", addService)
	router.PUT("/services/:id", updateService)
	router.DELETE("/services/:id", deleteService)

	router.Run(":8080") // starts a http server on localhost:8080
}

type Service struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

var services = []Service{}

func getServices(c *gin.Context) {
	c.JSON(http.StatusOK, services)
}

func addService(c *gin.Context) {
	var newService Service

	if err := c.ShouldBindJSON(&newService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, s := range services {
		if newService.Id == s.Id {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("service with id %s already exists", newService.Id)})
			return
		}
	}
	services = append(services, newService)
	c.JSON(http.StatusCreated, newService)
}

func updateService(c *gin.Context) {
	var updatedService Service
	id := c.Param("id")

	if err := c.ShouldBindJSON(&updatedService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, s := range services {
		if s.Id == id {
			services[i].Name = updatedService.Name
			services[i].State = updatedService.State
			c.JSON(http.StatusOK, services[i])
		}
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "service not found"})

}

func deleteService(c *gin.Context) {
	id := c.Param("id")

	for i, s := range services {
		if s.Id == id {
			services = append(services[:i], services[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "service deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("service with id %s not found", id)})
}
