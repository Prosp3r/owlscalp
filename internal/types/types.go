package types

import "time"

type Candle struct {
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Volume    float64 `json:"volume"`
}

type TradeSignal struct {
	ID                  string    `json:"id"`
	Symbol              string    `json:"symbol"`
	Direction           string    `json:"direction"`
	Entry               float64   `json:"entry"`
	Target              float64   `json:"target"`
	StopLoss            float64   `json:"stop_loss"`
	Leverage            int       `json:"leverage"`
	ExpectedGain        float64   `json:"expected_gain"`
	Timestamp           time.Time `json:"timestamp"`
	Tier                string    `json:"tier"`
	Reason              string    `json:"reason"`
	Confidence          float64   `json:"confidence"`
	Category            string    `json:"category"`
	Change24h           float64   `json:"change_24h"`
	RSI                 float64   `json:"rsi"`
	PositionSize        float64   `json:"position_size"`
	MaxHoldMinutes      int       `json:"max_hold_minutes"`
	TownCrierRSI        float64   `json:"towncrier_rsi"`
	TownCrierScore      int       `json:"towncrier_score"`
	TownCrierVolume     float64   `json:"towncrier_volume"`
	Resistance          float64   `json:"resistance"`
	DistanceToResist    float64   `json:"distance_to_resist"`
	SpeedMatrixTier     string    `json:"speed_matrix_tier"`
	VolumeProfileStrong bool      `json:"volume_profile_strong"`
	BuySellRatio        float64   `json:"buy_sell_ratio"`
	OrderBookImbalance  float64   `json:"order_book_imbalance"`
	EntryStrategy       string    `json:"entry_strategy"`
	ConfidenceScore     float64   `json:"confidence_score"`
	VolumeProfileRatio  float64   `json:"volume_profile_ratio"`
	OIDivergenceAlert   bool      `json:"oi_divergence_alert"`
	FundingSqueezeRisk  bool      `json:"funding_squeeze_risk"`
	MagnetEntry         bool      `json:"magnet_entry"`
}

type BacktestResult struct {
	TotalTrades  int
	WinRate      float64
	ProfitFactor float64
	TotalPnL     float64
	MaxDrawdown  float64
}
