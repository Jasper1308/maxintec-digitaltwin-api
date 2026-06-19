package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"maxintec-digitaltwin-api/internal/config"
	"maxintec-digitaltwin-api/internal/ordemservico"
	ordemMssql "maxintec-digitaltwin-api/internal/ordemservico/mssql"
	"maxintec-digitaltwin-api/internal/pessoa"
	pessoaMssql "maxintec-digitaltwin-api/internal/pessoa/mssql"
	"maxintec-digitaltwin-api/internal/platform/database"
	"maxintec-digitaltwin-api/internal/rastreador"
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

	rastreadorClient := rastreador.NewClientFromEnv()
	trackingCache := rastreador.NewMemoryCache()

	trackerWorker := rastreador.NewWorker(rastreadorClient, trackingCache)
	trackerWorker.Start(context.Background(), 1*time.Minute)

	osRepo := ordemMssql.NewRepository(db)
	osService := ordemservico.NewService(osRepo)
	osHandler := ordemservico.NewHandler(osService)

	pessoaRepo := pessoaMssql.NewRepository(db)
	pessoaService := pessoa.NewService(pessoaRepo)
	pessoaHandler := pessoa.NewHandler(pessoaService)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	v1 := r.Group("/api")
	{
		osHandler.RegisterRoutes(v1)
		pessoaHandler.RegisterRoutes(v1)

		v1.GET("/veiculos/posicoes", func(c *gin.Context) {
			posicoes := trackingCache.GetAll()
			
			if len(posicoes) == 0 {
				c.JSON(http.StatusOK, gin.H{
					"status":  "OK",
					"message": "Aguardando primeira coleta de dados do rastreador...",
					"data":    []interface{}{},
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
				"data":   posicoes,
			})
		})
	}

	fmt.Println("Servidor HTTP rodando na porta :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar o servidor HTTP: ", err)
	}
}