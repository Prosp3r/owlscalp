package indicators

import (
	"fmt"
	"time"

	"github.com/Prosp3r/owlscalp/internal/types"
)

func GenerateSignal(symbol string, candles []types.Candle) *types.TradeSignal {
	if len(candles) < 50 {
		return nil
	}

	lastCandle := candles[len(candles)-1]
	price := lastCandle.Close

	if len(candles)%50 == 0 {
		return &types.TradeSignal{
			ID:                  fmt.Sprintf("demo-%d", len(candles)),
			Symbol:              symbol,
			Direction:           "LONG",
			Entry:               price,
			Target:              price * 1.05,
			StopLoss:            price * 0.97,
			Leverage:            10,
			ExpectedGain:        5.0,
			Timestamp:           time.Now(),
			Tier:                "MEDIUM",
			Reason:              "Demo volume breakout",
			Confidence:          72.0,
			Category:            "MID_CAP",
			Change24h:           2.5,
			RSI:                 68.0,
			PositionSize:        100.0,
			MaxHoldMinutes:      45,
			TownCrierRSI:        67.0,
			TownCrierScore:      82,
			TownCrierVolume:     24500000,
			Resistance:          price * 1.08,
			DistanceToResist:    0.03,
			SpeedMatrixTier:     "MEDIUM",
			VolumeProfileStrong: true,
			BuySellRatio:        1.35,
			OrderBookImbalance:  0.42,
			EntryStrategy:       "AGGRESSIVE",
			ConfidenceScore:     72.0,
			VolumeProfileRatio:  3.8,
			OIDivergenceAlert:   false,
			FundingSqueezeRisk:  false,
			MagnetEntry:         false,
		}
	}

	if len(candles)%75 == 0 {
		return &types.TradeSignal{
			ID:                  fmt.Sprintf("demo-short-%d", len(candles)),
			Symbol:              symbol,
			Direction:           "SHORT",
			Entry:               price,
			Target:              price * 0.95,
			StopLoss:            price * 1.03,
			Leverage:            10,
			ExpectedGain:        5.0,
			Timestamp:           time.Now(),
			Tier:                "MEDIUM",
			Reason:              "Demo overbought rejection",
			Confidence:          68.0,
			Category:            "MID_CAP",
			Change24h:           -1.8,
			RSI:                 74.0,
			PositionSize:        100.0,
			MaxHoldMinutes:      45,
			TownCrierRSI:        73.0,
			TownCrierScore:      88,
			TownCrierVolume:     31200000,
			Resistance:          price * 1.02,
			DistanceToResist:    0.02,
			SpeedMatrixTier:     "MEDIUM",
			VolumeProfileStrong: true,
			BuySellRatio:        0.65,
			OrderBookImbalance:  -0.38,
			EntryStrategy:       "AGGRESSIVE",
			ConfidenceScore:     68.0,
			VolumeProfileRatio:  4.2,
			OIDivergenceAlert:   false,
			FundingSqueezeRisk:  false,
			MagnetEntry:         false,
		}
	}

	return nil
}
