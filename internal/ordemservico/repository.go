package ordemservico

import "context"

type Repository interface {
	GetTop5Recent(ctx context.Context) ([]OrdemServico, error)
	GetPainelAbertas(ctx context.Context) ([]PainelOrdemServico, error)
}