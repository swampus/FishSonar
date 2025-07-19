package rest

import (
    "encoding/json"
    "net/http"
    "strconv"
    "time"
    "github.com/swampus/fish-detector/internal/detector"
)

type Server struct {
    fish *detector.FishDetector
    mux  *http.ServeMux
}

func NewServer(fish *detector.FishDetector) *Server {
    s := &Server{fish: fish, mux: http.NewServeMux()}
    s.mux.HandleFunc("/api/check-fish", s.checkFish)
    s.mux.HandleFunc("/api/fish-history", s.fishHistoryHandler)
    s.mux.HandleFunc("/api/shark-advice", s.sharkAdviceHandler)
    return s
}

func (s *Server) Router() http.Handler { return s.mux }

func (s *Server) checkFish(w http.ResponseWriter, r *http.Request) {
    since := 10 * time.Second
    if v := r.URL.Query().Get("seconds"); v != "" {
        sec, err := strconv.Atoi(v)
        if err == nil && sec > 0 {
            since = time.Duration(sec) * time.Second
        }
    }
    res := s.fish.GetRecentFish(since)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}


func (s *Server) fishHistoryHandler(w http.ResponseWriter, r *http.Request) {
    minutes := 5 // default 5 минут
    if m := r.URL.Query().Get("minutes"); m != "" {
        if v, err := strconv.Atoi(m); err == nil && v > 0 && v <= 60 {
            minutes = v
        }
    }
    res := s.fish.GetRecentFish(time.Duration(minutes) * time.Minute)
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(res)
}


func (s *Server) sharkAdviceHandler(w http.ResponseWriter, r *http.Request) {
    minutes := 3
    if m := r.URL.Query().Get("minutes"); m != "" {
        if v, err := strconv.Atoi(m); err == nil && v > 0 && v <= 60 {
            minutes = v
        }
    }
    recent := s.fish.GetRecentFish(time.Duration(minutes) * time.Minute)
    var thiccCount, levCount, total int
    for _, evt := range recent {
        switch evt.Type {
        case detector.ThiccFish:
            thiccCount++
        case detector.LeverageFish:
            levCount++
        }
        total++
    }
    verdict := "Water is calm. Wait for more fish."
    if thiccCount > 0 || levCount > 0 || total >= 5 {
        verdict = "Big fish detected! Shark mode: ON."
    }
    resp := struct {
        Verdict      string `json:"verdict"`
        SharkMessage  string `json:"shark_message"`
        ThiccCount   int    `json:"thicc_fish"`
        LeverageCount int   `json:"leverage_fish"`
        TotalFish    int    `json:"total_fish"`
        TimeFrameMin int    `json:"timeframe_min"`
    }{
        Verdict:      verdict,
        SharkMessage: detector.RandomSharkAdvice(),
        ThiccCount:   thiccCount,
        LeverageCount: levCount,
        TotalFish:    total,
        TimeFrameMin: minutes,
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(resp)
}