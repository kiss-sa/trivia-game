package main

import (
	"log"
	"net/http"

	"trivia-game/backend/api"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting up...")

	boards, err := api.LoadBoardsFromCSV("data/boards.csv")
	if err != nil {
		log.Fatalf("Failed to load boards: %v", err)
	}
	log.Printf("Loaded %d boards", len(boards))

	r := gin.Default()
	r.SetTrustedProxies(nil)

	v1 := r.Group("/api")

	v1.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hi :)",
		})
	})

	v1.GET("/boards", api.GetBoards(boards))

	r.Run()
}
