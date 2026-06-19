package ordemservico

import "context"

type Service interface {
	GetRecentOrders(ctx context.Context) ([]OrdemServico, error)
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