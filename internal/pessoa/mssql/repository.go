package mssql

import (
	"context"
	"database/sql"
	"maxintec-digitaltwin-api/internal/pessoa"
)

type mssqlRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) pessoa.Repository {
	return &mssqlRepository{db: db}
}

func (r *mssqlRepository) GetTop5WithCNPJ(ctx context.Context) ([]pessoa.Pessoa, error) {
	// Importante: WITH (NOLOCK) para evitar locks em tabelas de produção durante leituras simples
	query := `SELECT TOP 5 Id, RazaoSocial, CNPJ FROM Pessoa WITH (NOLOCK) WHERE CNPJ IS NOT NULL`
	
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pessoas []pessoa.Pessoa
	for rows.Next() {
		var p pessoa.Pessoa
		if err := rows.Scan(&p.ID, &p.RazaoSocial, &p.CNPJ); err != nil {
			return nil, err
		}
		pessoas = append(pessoas, p)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pessoas, nil
}