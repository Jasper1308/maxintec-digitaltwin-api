package ordemservico

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
	router.GET("/ordens-servico/recentes", h.GetRecent)
	router.GET("/ordens-servico/painel", h.GetPainel)
}

func (h *Handler) GetRecent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	ordens, err := h.service.GetRecentOrders(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar ordens de serviço"})
		return
	}

	if ordens == nil {
		ordens = []OrdemServico{}
	}

	c.JSON(http.StatusOK, ordens)
}

func (h *Handler) GetPainel(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	dadosPainel, err := h.service.GetDashboard(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao gerar dados do painel"})
		return
	}

	if dadosPainel == nil {
		dadosPainel = []PainelOrdemServico{}
	}

	c.JSON(http.StatusOK, dadosPainel)
}