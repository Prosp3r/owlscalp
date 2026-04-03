package backtester

import (
	"context"

	"github.com/Prosp3r/0xwire/owlscalp/internal/filters"
	"github.com/Prosp3r/0xwire/owlscalp/internal/indicators"
	"github.com/Prosp3r/0xwire/owlscalp/internal/types"
)

type Backtester struct {
	initialBalance float64
	tradeAmount    float64
	commission     float64
}

func New(initialBalance, tradeAmount, commission float64) *Backtester {
	return &Backtester{
		initialBalance: initialBalance,
		tradeAmount:    tradeAmount,
		commission:     commission,
	}
}

func (b *Backtester) Run(ctx context.Context, symbol string, candles []types.Candle) (*types.BacktestResult, error) {
	balance := b.initialBalance
	wins := 0
	totalTrades := 0
	totalPnL := 0.0
	maxDrawdown := 0.0
	peak := balance

	for i := 50; i < len(candles)-10; i++ {
		window := candles[i-50 : i+1]

		signal := indicators.GenerateSignal(symbol, window)
		if signal == nil {
			continue
		}

		if !filters.ValidateSignal(signal, window) {
			continue
		}

		totalTrades++
		entry := signal.Entry
		exitPrice := 0.0
		maxHold := 30

		for j := i + 1; j < len(candles) && j < i+maxHold; j++ {
			c := candles[j]

			if signal.Direction == "LONG" {
				if c.High >= signal.Target {
					exitPrice = signal.Target
					break
				}
				if c.Low <= signal.StopLoss {
					exitPrice = signal.StopLoss
					break
				}
			} else {
				if c.Low <= signal.Target {
					exitPrice = signal.Target
					break
				}
				if c.High >= signal.StopLoss {
					exitPrice = signal.StopLoss
					break
				}
			}
			exitPrice = c.Close
		}

		if exitPrice == 0 {
			exitPrice = candles[i+1].Close
		}

		var pnl float64
		if signal.Direction == "LONG" {
			pnl = (exitPrice - entry) / entry * b.tradeAmount * float64(signal.Leverage)
		} else {
			pnl = (entry - exitPrice) / entry * b.tradeAmount * float64(signal.Leverage)
		}
		pnl -= b.commission * b.tradeAmount * 2

		balance += pnl
		totalPnL += pnl

		if pnl > 0 {
			wins++
		}

		if balance > peak {
			peak = balance
		}
		dd := (peak - balance) / peak * 100
		if dd > maxDrawdown {
			maxDrawdown = dd
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}

	winRate := 0.0
	if totalTrades > 0 {
		winRate = float64(wins) / float64(totalTrades) * 100
	}

	profitFactor := 0.0
	if totalPnL < 0 {
		profitFactor = totalPnL / (totalPnL * -1)
	} else if totalPnL > 0 {
		profitFactor = totalPnL / 0.0001
	}

	return &types.BacktestResult{
		TotalTrades:  totalTrades,
		WinRate:      winRate,
		ProfitFactor: profitFactor,
		TotalPnL:     totalPnL,
		MaxDrawdown:  maxDrawdown,
	}, nil
}
