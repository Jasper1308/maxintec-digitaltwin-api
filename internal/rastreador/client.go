package rastreador

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	username   string
	password   string
	appID      int
	token      string
}

func NewClientFromEnv() *Client {
	baseURL := os.Getenv("RASTREADOR_BASE_URL")
	username := os.Getenv("RASTREADOR_USERNAME")
	password := os.Getenv("RASTREADOR_PASSWORD")
	
	appID, err := strconv.Atoi(os.Getenv("RASTREADOR_APP_ID"))
	if err != nil {
		appID = 1718 
	}

	if baseURL == "" {
		baseURL = "http://maxintec.1gps.com.br:9870"
	}

	return &Client{
		httpClient: &http.Client{Timeout: 60 * time.Second},
		baseURL:    baseURL,
		username:   username,
		password:   password,
		appID:      appID,
	}
}

func (c *Client) Logon(ctx context.Context) error {
	body := LoginRequest{
		Username: c.username,
		Password: c.password,
		AppID:    c.appID,
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/seguranca/logon", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var res LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}

	if res.Status != "OK" {
		return fmt.Errorf("falha no logon do rastreador: %s", res.ResponseMessage)
	}

	c.token = res.Object.Token
	return nil
}

func (c *Client) FetchDadosNovos(ctx context.Context) ([]DadosNovos, error) {
	if c.token == "" {
		if err := c.Logon(ctx); err != nil {
			return nil, err
		}
	}

	url := fmt.Sprintf("%s/integracao/dados_novos", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro de conexao na rota de integracao: %w", err)
	}
	defer resp.Body.Close()

	var res IntegracaoResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON da fila: %w", err)
	}

	if res.Status == "VAZIO" {
		return []DadosNovos{}, nil
	}

	if res.Status != "OK" {
		return nil, fmt.Errorf("API recusou a integracao. Status: %s, Msg: %s", res.Status, res.ResponseMessage)
	}

	return res.Object, nil
}