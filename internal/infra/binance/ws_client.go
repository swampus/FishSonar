package binance

import (
    "encoding/json"
    "log"
    "strconv"
    "time"
    "github.com/gorilla/websocket"
    "github.com/swampus/fish-detector/internal/model"
)

type TradeHandler func(model.Trade)

type WSClient struct {
    handler TradeHandler
}

func NewWSClient(handler TradeHandler) *WSClient {
    return &WSClient{handler: handler}
}

func (c *WSClient) Start() {
    url := "wss://stream.binance.com:9443/ws/btcusdt@trade"
    conn, _, err := websocket.DefaultDialer.Dial(url, nil)
    if err != nil {
        log.Fatal("WebSocket dial error:", err)
    }
    defer conn.Close()
    log.Println("Connected to Binance WS:", url)

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("WS read error:", err)
            time.Sleep(2 * time.Second)
            continue
        }
        var data struct {
            Price        string `json:"p"`
            Quantity     string `json:"q"`
            Timestamp    int64  `json:"T"`
            IsBuyerMaker bool   `json:"m"`
            Symbol       string `json:"s"`
        }
        if err := json.Unmarshal(msg, &data); err != nil {
            log.Println("JSON decode error:", err)
            continue
        }
        price, _ := strconv.ParseFloat(data.Price, 64)
        qty, _ := strconv.ParseFloat(data.Quantity, 64)
        t := model.Trade{
            Symbol:       data.Symbol,
            Price:        price,
            Quantity:     qty,
            Timestamp:    data.Timestamp,
            IsBuyerMaker: data.IsBuyerMaker,
        }
        c.handler(t)
    }
}