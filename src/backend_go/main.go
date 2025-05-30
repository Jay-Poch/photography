package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type image struct {
	Id       string            `json:"id"`
	Location string            `json:"location"`
	Info     map[string]string `json:"info"`
}
type UserInput struct {
	Location             string            `json:"image_address"`
	Info                 map[string]string `json:"info"`
	Kameramodell         string            `json:"Kameramodell"`
	Verwendetes_Objektiv string            `json:"Verwendetes_Objektiv"`
	Belichtungszeit      string            `json:"Belichtungszeit"`
	Blende               string            `json:"Blende"`
	ISO                  string            `json:"ISO"`
	Beleuchtung          string            `json:"Beleuchtung"`
	Fokuspunkt           string            `json:"Fokuspunkt"`
	Fotogenre            string            `json:"Fotogenre"`
}

func main() {
	selecty("SELECT * FROM Photography WHERE id = '[value-1]'")
	images := map[string]image{
		"test": {Id: "test", Location: "myhome", Info: map[string]string{"test": "test"}},
		"test2": {Id: "test2", Location: "https://imgs.search.brave.com/K_opZj44YMhs8CF2dPWMiMkxY8RivnZ0E1FfN9HE-gM/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9pbWdz/LnNlYXJjaC5icmF2/ZS5jb20vR1JlMVlh/TkxKR2lvSWJOa2s5/ZEJGUkptcGIwOGJv/OTRvVkpCUkRWODZ6/dy9yczpmaXQ6NTAw/OjA6MDowL2c6Y2Uv/YUhSMGNITTZMeTlq/YjI1MC9aVzUwTFcx/aGJtRm5aVzFsL2Ju/UXRabWxzWlhNdVky/RnUvZG1FdVkyOXRM/Mk5rYmkxai9aMmt2/YVcxaFoyVXZaajFo/L2RYUnZMSEU5TnpB/dllUbGgvTlRNeVpU/UXRZVFZsTUMwMC9N/R1UwTFRrd09UTXRN/VEpoL09UWXlORGxq/TjJObUx6RXcvYzJs/dGNHeGxkMkY1YzNS/di9aVzVvWVc1alpY/bHZkWEpwL2JXRm5a/WE10VFc5eVpYSmwv/YzI5MWNtTmxjeko0/TG5CdS9adw", Info: map[string]string{
			"camera_model":  "Canon EOS R5",
			"lens_used":     "RF 50mm f/1.2L USM",
			"shutter_speed": "1/250 sec",
			"aperture":      "f/1.8",
			"ISO":           "200",
			"lighting":      "Studio Lampe",
			"focal_point":   "Beispiel Daten",
			"photo_genre":   "Portrait"}},
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
		c.JSON(http.StatusOK, images[id])
	})

	// POST /data → add a new item to the list
	router.POST("/add_image", func(c *gin.Context) {
		var input UserInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := uuid.New()
		postet := image{Id: id.String(), Location: input.Location, Info: map[string]string{
			"camera_model":  input.Kameramodell,
			"lens_used":     input.Verwendetes_Objektiv,
			"shutter_speed": input.Belichtungszeit,
			"aperture":      input.Blende,
			"ISO":           input.ISO,
			"lighting":      input.Beleuchtung,
			"focal_point":   input.Fokuspunkt,
			"photo_genre":   input.Fotogenre}}
		c.JSON(http.StatusOK, gin.H{
			"id": postet.Id})
		images[postet.Id] = postet

	})

	router.Run(":5000")
}
