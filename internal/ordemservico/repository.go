package ordemservico

import "context"

type Repository interface {
	GetTop5Recent(ctx context.Context) ([]OrdemServico, error)
}