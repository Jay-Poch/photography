package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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
		"test":  {Id: "test", Location: "myhome", Info: map[string]string{"test": "test"}},
		"test2": {Id: "test2", Location: "myhome", Info: map[string]string{"test": "test"}},
	}

	// Create a new web server
	router := gin.Default()
	// Enable CORS with custom settings
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:2020"}, // <-- frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// GET /data → return the list of items
	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, images[id].Info)
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

	router.Run(":5000")
}
