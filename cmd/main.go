package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"maxintec-digitaltwin-api/internal/config"
	ordemMssql "maxintec-digitaltwin-api/internal/ordemservico/mssql"
	pessoaMssql "maxintec-digitaltwin-api/internal/pessoa/mssql"
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

	pessoaRepo := pessoaMssql.NewRepository(db)
	osRepo := ordemMssql.NewRepository(db)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pessoas, err := pessoaRepo.GetTop5WithCNPJ(ctx)
	if err != nil {
		log.Printf("[Erro]: Falha ao buscar pessoas: %v\n", err)
	} else {
		fmt.Println("\nPessoas encontradas:")
		for _, p := range pessoas {
			fmt.Printf("ID: %d, Razão Social: %s, CNPJ: %s\n", p.ID, p.RazaoSocial, p.CNPJ)
		}
	}

	ordens, err := osRepo.GetTop5Recent(ctx)
	if err != nil {
		log.Printf("[Erro]: Falha ao buscar ordens de serviço: %v\n", err)
	} else {
		fmt.Println("\nOrdens de Serviço encontradas:")
		for _, o := range ordens {
			status := "Aberta"
			if o.DataHoraConclusao.Valid {
				status = fmt.Sprintf("Concluída em %s", o.DataHoraConclusao.Time.Format("02/01/2006 15:04"))
			}
			fmt.Printf("-> ID: %d, Número: %s, Cliente: %s, Status: %s\n", o.ID, o.Numero, o.RazaoSocial, status)
		}
		if len(ordens) == 0 {
			fmt.Println("[Aviso]: O SQL Server não devolveu nenhuma linha para a tabela dbo.OrdemServico.")
		}
	}
}