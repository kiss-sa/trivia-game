package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBoards(boards []*Board) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, boards)
	}
}
