package main

import (
    "log"
    "net/http"
    "github.com/swampus/fish-detector/internal/model"
    "github.com/swampus/fish-detector/internal/infra/binance"
    "github.com/swampus/fish-detector/internal/infra/rest"
    "github.com/swampus/fish-detector/internal/detector"
)

func main() {
    fishDetector := detector.NewFishDetector()

    // Binance WebSocket client
    ws := binance.NewWSClient(func(trade model.Trade) {
        fishDetector.ProcessTrade(trade)
    })
    go ws.Start()

    srv := rest.NewServer(fishDetector)
    log.Println("🐟 Fish Detector REST API running at :8080")
    log.Fatal(http.ListenAndServe(":8080", srv.Router()))
}