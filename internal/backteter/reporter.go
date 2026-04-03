package backtester

import (
	"fmt"

	"github.com/Prosp3r/owlscalp/internal/types"
)

func PrintResult(r *types.BacktestResult) {
	fmt.Printf("\n=== BACKTEST RESULTS ===\n")
	fmt.Printf("Total Trades:   %d\n", r.TotalTrades)
	fmt.Printf("Win Rate:       %.2f%%\n", r.WinRate)
	fmt.Printf("Total PnL:      $%.2f\n", r.TotalPnL)
	fmt.Printf("Profit Factor:  %.2f\n", r.ProfitFactor)
	fmt.Printf("Max Drawdown:   %.2f%%\n", r.MaxDrawdown)
	fmt.Println("========================\n")
}
