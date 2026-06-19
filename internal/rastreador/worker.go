package rastreador

import (
	"context"
	"log"
	"time"
)

type Worker struct {
	client *Client
	cache  *MemoryCache
}

func NewWorker(client *Client, cache *MemoryCache) *Worker {
	return &Worker{
		client: client,
		cache:  cache,
	}
}

func (w *Worker) Start(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	log.Println("[Worker] Monitoramento da fila de integração ativo...")

	w.execute(ctx)

	go func() {
		for {
			select {
			case <-ticker.C:
				w.execute(ctx)
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}

func (w *Worker) execute(ctx context.Context) {
	dados, err := w.client.FetchDadosNovos(ctx)
	if err != nil {
		log.Printf("[Worker] Erro na fila de integracao: %v\n", err)
		return
	}

	if len(dados) == 0 {
		log.Println("[Worker] Fila limpa. Aguardando novas transmissões dos rastreadores...")
		return
	}

	carrosValidos := map[int]string{
		297918: "RDW1C69",
		297598: "QPF1G05",
		296776: "AZB2J95",
		297120: "ASH3701",
	}

	count := 0
	for _, d := range dados {
		if placa, ok := carrosValidos[d.VeiculoID]; ok {
			if d.Latitude == 0 && d.Longitude == 0 {
				continue
			}

			var dataFormatada string
			if d.DataGps > 0 {
				tm := time.UnixMilli(d.DataGps)
				dataFormatada = tm.Format("02/01/2006 15:04:05")
			} else {
				dataFormatada = time.Now().Format("02/01/2006 15:04:05")
			}

			w.cache.Set(d.VeiculoID, LocalizacaoCarro{
				Placa:         placa,
				Latitude:      d.Latitude,
				Longitude:     d.Longitude,
				Velocidade:    d.Velocidade,
				Ligado:        d.Ignicao,
				UltimaLeitura: dataFormatada,
			})
			count++
		}
	}

	if count > 0 {
		log.Printf("[Worker] Sucesso! %d novas posições reais processadas e salvas no cache.\n", count)
	}
}