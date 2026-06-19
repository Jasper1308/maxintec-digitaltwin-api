package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"maxintec-digitaltwin-api/internal/config"
	"maxintec-digitaltwin-api/internal/ordemservico"
	ordemMssql "maxintec-digitaltwin-api/internal/ordemservico/mssql"
	"maxintec-digitaltwin-api/internal/platform/database"
)

func main() {
	cfg := config.Load()

	fmt.Println("Tentando conexão com SQL Server...")
	db, err := database.NewMSSQLConnection(cfg)
	if err != nil {
		log.Fatal("Erro crítico na conexão com o banco: ", err.Error())
	}
	defer db.Close()
	fmt.Println("Conexão com SQL Server estabelecida com sucesso!")

	osRepo := ordemMssql.NewRepository(db)
	osService := ordemservico.NewService(osRepo)
	osHandler := ordemservico.NewHandler(osService)

	// Usar gin.SetMode(gin.ReleaseMode) em produção real
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	v1 := r.Group("/api/v1")
	{
		osHandler.RegisterRoutes(v1)
	}

	fmt.Println("Servidor HTTP rodando na porta :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar o servidor HTTP: ", err)
	}
}