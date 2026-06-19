package ordemservico

import (
	"database/sql"
	"time"
)

type OrdemServico struct {
	ID                int
	Numero            string
	RazaoSocial       string
	Abertura          time.Time
	Prazo             time.Time
	DataHoraConclusao sql.NullTime
}