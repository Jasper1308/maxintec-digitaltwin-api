package mssql

import (
	"context"
	"database/sql"
	"maxintec-digitaltwin-api/internal/ordemservico"
)

type mssqlRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ordemservico.Repository {
	return &mssqlRepository{db: db}
}

func (r *mssqlRepository) GetTop5Recent(ctx context.Context) ([]ordemservico.OrdemServico, error) {
	// Incluído WITH (NOLOCK) para garantir concorrência limpa no banco de produção
	query := `
		SELECT TOP 5 Id, Numero, RazaoSocial, Abertura, Prazo, DataHoraConclusao 
		FROM dbo.OrdemServico WITH (NOLOCK)
		WHERE Numero IS NOT NULL AND Abertura IS NOT NULL
		ORDER BY Abertura DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ordens []ordemservico.OrdemServico
	for rows.Next() {
		var o ordemservico.OrdemServico
		err := rows.Scan(&o.ID, &o.Numero, &o.RazaoSocial, &o.Abertura, &o.Prazo, &o.DataHoraConclusao)
		if err != nil {
			return nil, err
		}
		ordens = append(ordens, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ordens, nil
}

func (r *mssqlRepository) GetPainelAbertas(ctx context.Context) ([]ordemservico.PainelOrdemServico, error) {
	query := `
		SELECT 
			Id, 
			ISNULL(Numero, ''), 
			ISNULL(RazaoSocial, 'Cliente Não Identificado'), 
			ResponsavelId, 
			Abertura, 
			Prazo,
			ISNULL(EnderecoRua, ''), 
			ISNULL(CONVERT(VARCHAR, EnderecoNumero), ''), 
			ISNULL(EnderecoBairro, ''), 
			ISNULL(EnderecoCep, '')
		FROM dbo.OrdemServico WITH (NOLOCK)
		WHERE DataHoraConclusao IS NULL 
		  AND Numero IS NOT NULL
		ORDER BY Abertura DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var painel []ordemservico.PainelOrdemServico
	for rows.Next() {
		var res ordemservico.PainelOrdemServico
		
		err := rows.Scan(
			&res.ID, 
			&res.Numero, 
			&res.Cliente, 
			&res.ResponsavelID, 
			&res.DataAbertura, 
			&res.Prazo,
			&res.Rua, 
			&res.Numero, 
			&res.Bairro, 
			&res.CEP,
		)
		if err != nil {
			return nil, err
		}
		painel = append(painel, res)
	}

	return painel, nil
}