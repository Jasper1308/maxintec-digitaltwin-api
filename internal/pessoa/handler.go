package pessoa

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/clientes/ativos", h.GetActive)
}

func (h *Handler) GetActive(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	clientes, err := h.service.GetActiveClients(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar clientes ativos"})
		return
	}

	if clientes == nil {
		clientes = []Pessoa{}
	}

	c.JSON(http.StatusOK, clientes)
}