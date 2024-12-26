package main

import (
	"github.com/Joshua1227/highwayLyricsApi/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/getAllTitles", handlers.GetAllTitles)
	router.GET("/getSongId/:id", handlers.GetSongById)
	router.GET("/searchSongs/:key", handlers.SearchSongs)

	router.Run("localhost:8080")
}
