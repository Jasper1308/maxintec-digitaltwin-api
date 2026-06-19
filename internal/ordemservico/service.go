package ordemservico

import (
	"context"
	"fmt"
	"time"
)

type Service interface {
	GetRecentOrders(ctx context.Context) ([]OrdemServico, error)
	GetDashboard(ctx context.Context) ([]PainelOrdemServico, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetRecentOrders(ctx context.Context) ([]OrdemServico, error) {
	return s.repo.GetTop5Recent(ctx)
}

func (s *service) GetDashboard(ctx context.Context) ([]PainelOrdemServico, error) {
	ordens, err := s.repo.GetPainelAbertas(ctx)
	if err != nil {
		return nil, err
	}
	
	agora := time.Now()
	for i := range ordens {
		duracao := agora.Sub(ordens[i].DataAbertura)
		
		dias := int(duracao.Hours() / 24)
		horas := int(duracao.Hours()) % 24

		if dias > 0 {
			ordens[i].TempoDecorrido = fmt.Sprintf("%d dia(s) e %d hora(s)", dias, horas)
		} else {
			ordens[i].TempoDecorrido = fmt.Sprintf("%d hora(s)", horas)
		}
	}

	return ordens, nil
}