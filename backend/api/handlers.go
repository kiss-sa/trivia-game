package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	store BoardStore
}

func NewHandler(store BoardStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) GetBoards(c *gin.Context) {
	boards, err := h.store.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, boards)
}

func (h *Handler) GetBoard(c *gin.Context) {
	board, err := h.store.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, board)
}
