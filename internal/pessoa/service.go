package pessoa

import "context"

type Service interface {
	GetActiveClients(ctx context.Context) ([]Pessoa, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetActiveClients(ctx context.Context) ([]Pessoa, error) {
	return s.repo.GetClientesComOSAberta(ctx)
}