package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define the structure of your data (like a dictionary in Python)
type Item struct {
	Name string `json:"name"` // Expect JSON with a "name" field
}

// This will store our data in memory
type image struct {
	Id       string            `json:"id"`
	Location string            `json:"location"`
	Info     map[string]string `json:"info"`
}

func main() {
	images := map[string]image{
		"test":  image{Id: "test", Location: "myhome", Info: map[string]string{"test": "test"}},
		"test2": image{Id: "test2", Location: "myhome", Info: map[string]string{"test": "test"}},
	}

	// Create a new web server
	router := gin.Default()

	// GET /data → return the list of items
	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, images[id])
	})

	// POST /data → add a new item to the list
	router.POST("/data", func(c *gin.Context) {
		var newItem Item

		// Try to read JSON from the request
		if err := c.BindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Add the item to the data store
		//dataStore = append(dataStore, newItem)
		c.JSON(http.StatusCreated, gin.H{"message": "Item added!"})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
