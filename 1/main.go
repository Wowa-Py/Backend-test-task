package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"github.com/joho/godotenv"
	"log"
	"os"
)

err := godotenv.Load()
if err != nil {
	log.Fatal("Ошибка загрузки файла .env")
}

baseURL := os.Getenv("BASE_URL")

type CoinData struct {
	ID                    string      `json:"id"`
	Symbol                string      `json:"symbol"`
	Name                  string      `json:"name"`
	Image                 string      `json:"image"`
	CurrentPrice          float64     `json:"current_price"`
	MarketCap             float64     `json:"market_cap"`
	MarketCapRank         int         `json:"market_cap_rank"`
	FullyDilutedValue     float64     `json:"fully_diluted_valuation"`
	TotalVolume           float64     `json:"total_volume"`
	High24h               float64     `json:"high_24h"`
	Low24h                float64     `json:"low_24h"`
	PriceChange24h        float64     `json:"price_change_24h"`
	PriceChangePct24h     float64     `json:"price_change_percentage_24h"`
	MarketCapChange24h    float64     `json:"market_cap_change_24h"`
	MarketCapChangePct24h float64     `json:"market_cap_change_percentage_24h"`
	CirculatingSupply     float64     `json:"circulating_supply"`
	TotalSupply           float64     `json:"total_supply"`
	MaxSupply             float64     `json:"max_supply"`
	Ath                   float64     `json:"ath"`
	AthChangePct          float64     `json:"ath_change_percentage"`
	AthDate               string      `json:"ath_date"`
	Atl                   float64     `json:"atl"`
	AtlChangePct          float64     `json:"atl_change_percentage"`
	AtlDate               string      `json:"atl_date"`
	Roi                   interface{} `json:"roi"`
	LastUpdated           string      `json:"last_updated"`
}

type CoinGeckoClient struct {
	BaseURL     string
	Cache       map[string]CoinData
	LastUpdated int64
}

func NewCoinGeckoClient() *CoinGeckoClient {
	return &CoinGeckoClient{
		BaseURL:     baseURL,
		Cache:       make(map[string]CoinData),
		LastUpdated: 0,
	}
}

func (c *CoinGeckoClient) GetCoinData(coinID string) (*CoinData, error) {
	if time.Now().Unix()-c.LastUpdated > 600 {
		err := c.UpdateCache()
		if err != nil {
			return nil, err
		}
	}

	coinData, ok := c.Cache[coinID]
	if ok {
		return &coinData, nil
	} else {
		return nil, fmt.Errorf("Coin with ID %s not found", coinID)
	}
}

func (c *CoinGeckoClient) UpdateCache() error {
	url := fmt.Sprintf("%s/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1", c.BaseURL)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var coinData []CoinData
	err = json.Unmarshal(body, &coinData)
	if err != nil {
		return err
	}

	c.Cache = make(map[string]CoinData)
	for _, coin := range coinData {
		c.Cache[coin.ID] = coin
	}

	c.LastUpdated = time.Now().Unix()

	return nil
}

func main() {
	client := NewCoinGeckoClient()

	// Пример использования (возможность получать курс определенной криптовалюты)
	coinID := "bitcoin"
	coinData, err := client.GetCoinData(coinID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Текущая цена %s: %.2f USD\n", coinData.Name, coinData.CurrentPrice)
}
