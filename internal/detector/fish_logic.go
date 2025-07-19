package detector

import (
    "fmt"
    "sync"
    "time"
    "sort"
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

type FishDetector struct {
    mu     sync.Mutex
    fish   []FishEvent
    trades []model.Trade
}

const maxFish = 10000

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

func classifyFish(trade model.Trade, recent []model.Trade) (FishType, string) {
    med := medianQuantity(recent)

    //bots or not interesting fish
    if trade.Quantity < 0.066 {
        return "", ""
    }

    // 1. Big Fish — X10 medians to big
    if trade.Quantity > med*20 && trade.Quantity > 1 && med > 0 {
         return ThiccFish, "ANOMALY: Really huge fish!"
    }

    // 2. YOLO Fish
    if trade.Leverage >= 20 && trade.Quantity > med*2 {
        return LeverageFish, "YOLO Fish: All-in with crazy leverage!"
    }

    // 3. Pump Fish
    if priceJumped(trade, recent, 0.02) && trade.Quantity > med*3 {
        return ThiccFish, "PUMP FISH: Price spiked up and big volume!"
    }

    // 4. Dump Fish
    if priceDropped(trade, recent, 0.02) && trade.Quantity > med*3 {
        return DumbFish, "DUMP FISH: Price dumped and big volume!"
    }

    // 5. Whale at Night
    if isNightTime(trade.Timestamp) && trade.Quantity > med*5 {
        return SleepyFish, "NIGHT WHALE: Big fish swims while all are sleeping."
    }

    // 6. High Roller —
    high, low := highLow(recent, 600) // за 10 минут
    if high > 0 && (trade.Price > high*0.99 || trade.Price < low*1.01) && trade.Quantity > med*3 {
        return ThiccFish, "High Roller: Big volume at market extreme!"
    }

    // 7. Cluster Fish —
    if clusterVolume(recent, 10, med*30) && med > 0.05 {
        return NormieFish, "Cluster Fish: Big trade burst."
    }

    // norm fish
    if trade.Quantity >= med && med > 0 {
        return NormieFish, "Just an average fish swimming by."
    }

    return "", ""
}

func isNightTime(ts int64) bool {
    t := time.UnixMilli(ts).UTC()
    return t.Hour() >= 0 && t.Hour() <= 5
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

func highLow(trades []model.Trade, periodSec int64) (high, low float64) {
    high, low = 0.0, 1e18
    cutoff := time.Now().Add(-time.Duration(periodSec) * time.Second).UnixMilli()
    for _, t := range trades {
        if t.Timestamp >= cutoff {
            if t.Price > high {
                high = t.Price
            }
            if t.Price < low {
                low = t.Price
            }
        }
    }
    return
}

func medianQuantity(trades []model.Trade) float64 {
    n := len(trades)
    if n == 0 {
        return 0
    }
    arr := make([]float64, n)
    for i, t := range trades {
        arr[i] = t.Quantity
    }
    sort.Float64s(arr)
    mid := n / 2
    if n%2 == 0 {
        return (arr[mid-1] + arr[mid]) / 2
    }
    return arr[mid]
}

func averageQuantity(trades []model.Trade) float64 {
    if len(trades) == 0 {
        return 0
    }
    sum := 0.0
    for _, t := range trades {
        sum += t.Quantity
    }
    return sum / float64(len(trades))
}

func (d *FishDetector) ProcessTrade(trade model.Trade) {
    d.mu.Lock()
    defer d.mu.Unlock()

    d.trades = append(d.trades, trade)
    if len(d.trades) > 1000 {
        d.trades = d.trades[1:]
    }

    fish, msg := classifyFish(trade, d.trades)
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

func priceJumped(trade model.Trade, recent []model.Trade, threshold float64) bool {
    cutoff := time.Now().Add(-5 * time.Minute).UnixMilli()
    minPrice := trade.Price
    for _, t := range recent {
        if t.Timestamp >= cutoff && t.Price < minPrice {
            minPrice = t.Price
        }
    }
    return (trade.Price-minPrice)/minPrice > threshold
}

func priceDropped(trade model.Trade, recent []model.Trade, threshold float64) bool {
    cutoff := time.Now().Add(-5 * time.Minute).UnixMilli()
    maxPrice := trade.Price
    for _, t := range recent {
        if t.Timestamp >= cutoff && t.Price > maxPrice {
            maxPrice = t.Price
        }
    }
    return (maxPrice-trade.Price)/maxPrice > threshold
}

func clusterVolume(recent []model.Trade, sec int, volumeThreshold float64) bool {
    cutoff := time.Now().Add(-time.Duration(sec) * time.Second).UnixMilli()
    total := 0.0
    for _, t := range recent {
        if t.Timestamp >= cutoff {
            total += t.Quantity
        }
    }
    return total > volumeThreshold
}