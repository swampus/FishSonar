package detector

import (
    "fmt"
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}


var FishMessages = map[FishType][]string{
    ThiccFish: {
        "Big splash! That's a lot of fins in the water.",
        "Someone just moved the ocean.",
        "Thicc fish sighted — brace yourselves!",
        "Heavy swimmer! Hope you know where you're going.",
        "That's not just a fish, that's almost a submarine.",
        "That’s enough volume to make the kraken jealous.",
        "Making waves that reach the shore.",
        "Biggest fish since the last exchange maintenance.",
        "That trade just woke up the whales.",
        "Splash so big, we need a mop.",
        "Trading like there’s no tomorrow.",
        "When you can't decide — so you buy the whole sea.",
        "If that's not a statement, what is?",
    },
    SleepyFish: {
        "Night voyage: fish rarely sleep.",
        "Some fish prefer the moonlight.",
        "Trading at night — respect for dedication!",
        "Fish is up while others are counting plankton.",
        "Silent water at night, but some fish keep moving.",
        "Trading at night? The sea is calm, but the mind is stormy.",
        "You know it's night when even bots are yawning.",
        "Fish swimming in the dark, maybe dreaming of profits.",
        "While most sleep, some chase shadows.",
        "Trading by moonlight — hope you packed your night goggles.",
        "At this hour? Hope you set a stop-loss *and* an alarm clock.",
        "Risky trades and night raids — classic combo.",
        "Is it insomnia or just crypto FOMO?",
    },
    DumbFish: {
        "Swimming against the current? Brave or bold?",
        "Some fish just love adventure — right into the net.",
        "Every journey needs a start — sometimes it starts at the top.",
        "Fish follows the shiny lure — classic move.",
        "Is that a wave or a leap of faith?",
        "Chasing bubbles since the dawn of time.",
        "Bought the top, sold the bottom — innovation!",
        "Making bold moves with questionable logic.",
        "Trust the plan... wait, what plan?",
        "Riding the wave — straight into the rocks.",
        "When in doubt, do the opposite of this fish.",
        "This fish thinks 'ATH' means 'Always Trade Here'.",
        "Knife-catching: Olympic edition.",
    },
    NormieFish: {
        "Just keep swimming.",
        "A regular day in the ocean.",
        "This fish is minding its own business.",
        "Sometimes you just need to swim for the sake of swimming.",
        "Not every swim is an adventure, but every fish counts.",
        "Just another day at the fish market.",
        "Swimming along, minding its own wallet.",
        "No drama, just fins and trades.",
        "If in doubt, just swim forward.",
        "Most fish just want a quiet swim.",
        "Not every trade is epic, but every fish is unique.",
        "In the ocean, even small ripples matter.",
        "No news is good news.",
    },
    LeverageFish: {
        "Some fish use fins, some use turbo fins.",
        "Swimming with extra floaties — leverage detected.",
        "High tide, high risk — leverage in the ocean.",
        "Hope those fins are strong enough for the current!",
        "Fish went for a ride with the wind behind.",
        "Leverage on? May Poseidon help you.",
        "Some fish like to swim close to the edge.",
        "Turbo-fins activated: full risk ahead!",
        "Leverage: because the ocean isn’t dangerous enough.",
        "Liquidation is just an oceanic adventure.",
        "If this fish swims any faster, it’ll hit the reef.",
        "‘YOLO’ isn’t an investment strategy... unless you’re a leverage fish.",
        "Double or nothing — the fish way.",
    },
}

var SharkMessages = []string{
    "Lots of fish detected! Shark mode recommended.",
    "Blood in the water? Now’s your chance.",
    "Time to hunt — the school is in session.",
    "When the sea is full of fish, the sharks never sleep.",
    "Smells like opportunity. Shark, are you ready?",
    "Fins everywhere — don’t miss your feeding frenzy.",
    "Alert: Waters crowded. Perfect time for a big bite.",
    "Advice: When you see fish, move like a shark!",
    "Opportunity detected: deploy the shark!",
    "Scent of fresh fish in the water — action time.",
    "If you wait too long, someone else gets the sushi.",
    "Sharks don’t sleep, and neither should you right now.",
    "Feeding frenzy mode: ON.",
}

func randomFishMessage(ft FishType) string {
    msgs := FishMessages[ft]
    if len(msgs) == 0 {
        return ""
    }
    return msgs[rand.Intn(len(msgs))]
}


func randomSharkAdvice() string {
    return SharkMessages[rand.Intn(len(SharkMessages))]
}

func formatFishMessage(ft FishType, symbol string, qty, price float64) string {
    msg := randomFishMessage(ft)
    return fmt.Sprintf("%s [%s] %.4f BTC @ $%.2f", msg, symbol, qty, price)
}

func RandomSharkAdvice() string {
    return SharkMessages[rand.Intn(len(SharkMessages))]
}