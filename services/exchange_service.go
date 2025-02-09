package services

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type ExchangeService struct {
    apiKey string
}

func NewExchangeService(apiKey string) *ExchangeService {
    return &ExchangeService{apiKey: apiKey}
}

type ExchangeResponse struct {
    Success bool               `json:"success"`
    Rates   map[string]float64 `json:"rates"`
}

func (s *ExchangeService) ConvertCurrency(amount float64, from, to string) (float64, error) {
    url := fmt.Sprintf("https://api.exchangerate-api.com/v4/latest/%s?access_key=%s", from, s.apiKey)
    
    resp, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    var result ExchangeResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return 0, err
    }

    if rate, ok := result.Rates[to]; ok {
        return amount * rate, nil
    }
    return 0, fmt.Errorf("currency conversion not available")
}