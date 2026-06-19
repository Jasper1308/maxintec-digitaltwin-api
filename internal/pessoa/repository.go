package pessoa

import "context"

type Repository interface {
	GetTop5WithCNPJ(ctx context.Context) ([]Pessoa, error)
}