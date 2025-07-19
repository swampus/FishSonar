# FishSonar — Real-Time Market Fish Detector

**FishSonar** is a professional, Clean Architecture Go service for real-time detection of "fishy" events on crypto exchanges.

Ever wanted to spot a THICC FISH or a SLEEPY FISH in your trade stream? This project analyzes trades in real time, connects to Binance, and exposes a REST API for you or your bots.

**Server**: stream.binance.com:9443/ws/btcusdt@trade

---

## Features

- **Real-time detection** of suspicious ("fishy") trades on Binance.
- Detects several types of market fish:  
  `THICC FISH`, `SLEEPY FISH`, `DUMB FISH`, `ORDINARY FISH`, `LEVERAGE FISH`
- **Clean, modular architecture** — easy to extend or adapt for your own experiments.
- REST API:
   - `/api/check-fish` — check the last trade for fishiness
   - `/api/fish-history` — view detected fish events
   - `/api/shark-advice` — receive questionable shark wisdom
- Configurable & testable (with unit tests)
- Designed for further AI/ML integration

---

# API Reference

## `GET /api/check-fish`

**Description:**  
Check for fishy trades in the recent time window.

**Query Parameters:**
- `seconds` (optional, integer):  
  Time window in **seconds** to scan for fishy trades.  
  Default: `10` (last 10 seconds).  
  Example: `?seconds=15`

**Example Request:**
```bash
curl http://localhost:8080/api/check-fish?seconds=15
```
``` example Response
[
  {
    "Trade": {
      "symbol": "BTCUSDT",
      "price": 117911.43,
      "quantity": 0.16,
      "timestamp": 1752956344955,
      "is_buyer_maker": true,
      "leverage": 0
    },
    "Type": "ORDINARY FISH",
    "Message": "This fish is minding its own business.",
    "Time": "2025-07-19T23:19:00.046506+03:00"
  },
  {
    "Trade": {
      "symbol": "BTCUSDT",
      "price": 117911.43,
      "quantity": 0.16,
      "timestamp": 1752956344955,
      "is_buyer_maker": true,
      "leverage": 0
    },
    "Type": "ORDINARY FISH",
    "Message": "Just another day at the fish market.",
    "Time": "2025-07-19T23:19:00.047506+03:00"
  }
]
```

## `GET /api/fish-history`

**Description:**  
Get history of detected fish events over a recent time period.

**Query Parameters:**
- `minutes` (optional, integer):  
  Time window in minutes to scan for fishy trades (1–60).  
  Default: `5` (last 5 minutes).  
  Example: `?minutes=20`

**Example Request:**
```bash
curl http://localhost:8080/api/fish-history?minutes=5
```

``` example response:
[
  {
    "Trade": {
      "symbol": "BTCUSDT",
      "price": 117950,
      "quantity": 0.21229,
      "timestamp": 1752956272115,
      "is_buyer_maker": true,
      "leverage": 0
    },
    "Type": "ORDINARY FISH",
    "Message": "In the ocean, even small ripples matter.",
    "Time": "2025-07-19T23:17:47.207506+03:00"
  }
]
```

## `GET /api/shark-advice`

**Description:**  
Get a shark verdict and a random shark trading message based on recent market activity.

```bash
curl http://localhost:8080/api/shark-advice
```

``` Example response
{
  "verdict": "Big fish detected! Shark mode: ON.",
  "shark_message": "Scent of fresh fish in the water — action time.",
  "thicc_fish": 0,
  "leverage_fish": 0,
  "total_fish": 13,
  "timeframe_min": 3
}
```


## Project Structure

- `cmd/` — main entry point
- `internal/detector/` — all fish detection logic
- `internal/infra/` — integrations (Binance WebSocket, REST API)
- `internal/model/` — domain models (Trade, FishEvent, etc.)
- `config/` — configuration files
- `scripts/` — automation and helpers

---

## Quick Start

1. **Install Go** (if not already):
    ```bash
    # For Linux/macOS:
    sudo apt install golang-go
    # or follow instructions: https://go.dev/doc/install
    ```

2. **Clone and build the project:**
    ```bash
    git clone https://github.com/swampus/FishSonar.git
    cd FishSonar
    go build ./cmd/main.go
    ./main
   
    go run ./cmd/main.go
    ```
