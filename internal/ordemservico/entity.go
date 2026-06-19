package ordemservico

import "time"

type OrdemServico struct {
	ID                int        `json:"id"`
	Numero            string     `json:"numero"`
	RazaoSocial       string     `json:"razao_social"`
	CNPJ              string     `json:"cnpj,omitempty"`
	Abertura          time.Time  `json:"data_abertura"`
	Prazo             time.Time  `json:"prazo"`
	DataHoraConclusao *time.Time `json:"data_hora_conclusao,omitempty"`
	Status            string     `json:"status"`
}

type PainelOrdemServico struct {
	ID             int        `json:"id"`
	Numero         string     `json:"numero"`
	Cliente        string     `json:"cliente"`
	ResponsavelID  *int       `json:"responsavel_id"` 
	DataAbertura   time.Time  `json:"data_abertura"`
	Prazo          time.Time  `json:"prazo"`
	TempoDecorrido string     `json:"tempo_decorrido"` 
	Rua       string `json:"rua"`
	NumeroEnd string `json:"numero_end"` 
	Bairro    string `json:"bairro"`
	CEP       string `json:"cep"`
}