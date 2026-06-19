package ordemservico

import "time"

type PainelOrdemServico struct {
	ID             int       `json:"id"`
	Numero         string    `json:"numero"`
	Cliente        string    `json:"cliente"`
	ResponsavelID  *int      `json:"responsavel_id"`
	DataAbertura   time.Time `json:"data_abertura"`
	Prazo          time.Time `json:"prazo"`
	TempoDecorrido string    `json:"tempo_decorrido"`
	Rua    string `json:"rua"`
	Numero string `json:"numero_end"`
	Bairro string `json:"bairro"`
	CEP    string `json:"cep"`
}