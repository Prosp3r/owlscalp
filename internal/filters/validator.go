package filters

import (
	"github.com/Prosp3r/owlscalp/internal/types"
)

func ValidateSignal(signal *types.TradeSignal, candles []types.Candle) bool {
	if signal == nil {
		return false
	}
	return true
}
