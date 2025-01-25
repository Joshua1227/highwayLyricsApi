package main

import (
	"github.com/Joshua1227/highwayLyricsApi/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/getAllTitles", handlers.GetAllTitles)
	router.GET("/getSongId/:id", handlers.GetSongById)
	router.GET("/searchSongs/:key", handlers.SearchSongs)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
