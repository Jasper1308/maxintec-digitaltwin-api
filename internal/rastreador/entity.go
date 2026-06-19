package rastreador

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AppID    int    `json:"appid"`
}

type LoginResponse struct {
	Status          string `json:"status"`
	ResponseMessage string `json:"responseMessage"`
	Object          struct {
		Token string `json:"token"`
	} `json:"object"`
}

type LocalizacaoCarro struct {
	Placa         string  `json:"placa"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Velocidade    float64 `json:"velocidade"`
	Ligado        bool    `json:"ligado"`
	UltimaLeitura string  `json:"ultima_leitura"`
}

type IntegracaoResponse struct {
	Status          string       `json:"status"`
	ResponseMessage string       `json:"responseMessage"`
	Object          []DadosNovos `json:"object"`
}

type DadosNovos struct {
	VeiculoID  int     `json:"veiculoId"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Velocidade float64 `json:"velocidade"`
	Ignicao    bool    `json:"ignicao"`
	DataGps    int64   `json:"dataGps"`
}