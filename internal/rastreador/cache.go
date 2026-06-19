package rastreador

import (
	"sync"
)

type MemoryCache struct {
	mu     sync.RWMutex
	carros map[int]LocalizacaoCarro
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		carros: make(map[int]LocalizacaoCarro),
	}
}

func (c *MemoryCache) Set(id int, novaLoc LocalizacaoCarro) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if novaLoc.Latitude == 0 && novaLoc.Longitude == 0 {
		if antigo, existe := c.carros[id]; existe {
			antigo.Ligado = novaLoc.Ligado
			antigo.Velocidade = 0
			c.carros[id] = antigo
			return
		}
	}

	c.carros[id] = novaLoc
}

func (c *MemoryCache) Get(id int) (LocalizacaoCarro, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	loc, ok := c.carros[id]
	return loc, ok
}

func (c *MemoryCache) GetAll() map[int]LocalizacaoCarro {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	copia := make(map[int]LocalizacaoCarro)
	for k, v := range c.carros {
		copia[k] = v
	}
	return copia
}