package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Entity
type album struct {
	ID     string  `json:"id`
	Title  string  `json:"title"`
	Artist string  `json:"artist""`
	Price  float64 `json:"price"`
}

// Mock Data
var albums = []album{
	{
		ID:     "1",
		Title:  "Blue Train",
		Artist: "John Coltrance",
		Price:  56.99,
	},
	{
		ID:     "2",
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99,
	},
	{
		ID:     "3",
		Title:  "Sarah Vaunhan and Clifford Brown",
		Artist: "Sarah Vaughan",
		Price:  36.99,
	},
}

// GET : get album list
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	// create empty album data
	var newAlbum album

	// If newAlbum is not able to bind Request Body, then return
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// append newAlbum data
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	// Initialize a Gin Router
	router := gin.Default()

	// HTTP Router
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")

}
