package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/Prosp3r/owlscalp/internal/backtester"
	"github.com/Prosp3r/owlscalp/internal/types"
	"github.com/Prosp3r/owlscalp/internal/utils"
)

func generateSyntheticCandles(count int) []types.Candle {
	rand.Seed(time.Now().UnixNano())
	candles := make([]types.Candle, count)
	price := 100.0

	for i := 0; i < count; i++ {
		price += rand.Float64()*0.8 - 0.4
		if price < 80 {
			price = 80
		}
		candles[i] = types.Candle{
			Timestamp: time.Now().Add(time.Duration(i) * time.Minute).UnixMilli(),
			Open:      price - rand.Float64()*0.3,
			High:      price + rand.Float64()*0.5,
			Low:       price - rand.Float64()*0.5,
			Close:     price,
			Volume:    1000000 + rand.Float64()*5000000,
		}
		price = candles[i].Close
	}
	return candles
}

func main() {
	utils.InitLogger()

	candles := generateSyntheticCandles(800)

	fmt.Println("Running Owlscalp backtester on 800 synthetic candles...")

	bt := backtester.New(1000, 100, 0.0005)

	ctx := context.Background()
	result, err := bt.Run(ctx, "BTCUSDT", candles)
	if err != nil {
		fmt.Println("Backtest error:", err)
		return
	}

	backtester.PrintResult(result)

	fmt.Println("Owlscalp backtester is now fully modular and ready.")
	fmt.Println("Next step: reply with exactly one of these:")
	fmt.Println("- Implement indicators folder")
	fmt.Println("- Implement filters folder")
	fmt.Println("- Add config loader")
}
