package model

type Trade struct {
    Symbol       string  `json:"symbol"`
    Price        float64 `json:"price"`
    Quantity     float64 `json:"quantity"`
    Timestamp    int64   `json:"timestamp"`
    IsBuyerMaker bool    `json:"is_buyer_maker"`
    Leverage     float64 `json:"leverage"`
}