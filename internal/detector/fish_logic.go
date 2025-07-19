package detector

import (
    "fmt"
    "sync"
    "time"
    "github.com/swampus/fish-detector/internal/model"
)

type FishType string

const (
    ThiccFish     FishType = "THICC FISH"
    SleepyFish    FishType = "SLEEPY FISH"
    DumbFish      FishType = "DUMB FISH"
    NormieFish    FishType = "ORDINARY FISH"
    LeverageFish  FishType = "LEVERAGE FISH"
)

type FishEvent struct {
    Trade   model.Trade
    Type    FishType
    Message string
    Time    time.Time
}

const maxFish = 10000 // сколько хотим хранить

type FishDetector struct {
    mu   sync.Mutex
    fish []FishEvent
}

func NewFishDetector() *FishDetector {
    return &FishDetector{
        fish: make([]FishEvent, 0, maxFish),
    }
}

func (d *FishDetector) addFish(evt FishEvent) {
    if len(d.fish) >= maxFish {
        d.fish = d.fish[1:] // удаляем самый старый
    }
    d.fish = append(d.fish, evt)
}

func (d *FishDetector) ProcessTrade(trade model.Trade) {
    d.mu.Lock()
    defer d.mu.Unlock()
    fish, msg := classifyFish(trade)
    if fish != "" {
        evt := FishEvent{
            Trade:   trade,
            Type:    fish,
            Message: msg,
            Time:    time.Now(),
        }
        d.addFish(evt) 
        logFish(evt)
    }
}

func (d *FishDetector) GetRecentFish(since time.Duration) []FishEvent {
    d.mu.Lock()
    defer d.mu.Unlock()
    res := make([]FishEvent, 0)
    cutoff := time.Now().Add(-since)
    for _, evt := range d.fish {
        if evt.Time.After(cutoff) {
            res = append(res, evt)
        }
    }
    return res
}

func classifyFish(trade model.Trade) (FishType, string) {
    switch {
    case trade.Leverage >= 2:
        return LeverageFish, randomFishMessage(LeverageFish)
    case trade.Quantity >= 10:
        return ThiccFish, randomFishMessage(ThiccFish)
    case trade.Quantity >= 0.1 && isNightTime(trade.Timestamp):
        return SleepyFish, randomFishMessage(SleepyFish)
    case trade.Quantity >= 0.1 && isNearHigh(trade.Price):
        return DumbFish, randomFishMessage(DumbFish)
    case trade.Quantity >= 0.1:
        return NormieFish, randomFishMessage(NormieFish)
    default:
        return "", ""
    }
}

func isNightTime(ts int64) bool {
    t := time.UnixMilli(ts).UTC()
    return t.Hour() >= 0 && t.Hour() <= 5
}

func isNearHigh(price float64) bool {
    // TODO: добавить аналитику high/low, если надо
    return false
}

func logFish(evt FishEvent) {
    fmt.Println("\n 🐟 ", evt.Type, "ALERT!")
    fmt.Println("=======================")
    fmt.Println("Symbol    :", evt.Trade.Symbol)
    fmt.Println("Price     :", evt.Trade.Price)
    fmt.Println("Quantity  :", evt.Trade.Quantity)
    fmt.Println("USD       :", evt.Trade.Price*evt.Trade.Quantity)
    fmt.Println("Time      :", time.UnixMilli(evt.Trade.Timestamp).UTC().Format("2006-01-02 15:04:05"))
    fmt.Println("Side      :", direction(evt.Trade.IsBuyerMaker))
    fmt.Println("Msg       :", evt.Message)
    fmt.Println("=======================\n")
}

func direction(isBuyerMaker bool) string {
    if isBuyerMaker {
        return "SELL"
    }
    return "BUY"
}