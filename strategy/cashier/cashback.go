package cashier

import "math"

type CashbackCashStrategy struct {
	condition float64
	back      float64
}

func (cashback *CashbackCashStrategy) CashAlgorithm(total float64) float64 {
	return total - (math.Floor(total/cashback.condition) * cashback.back)
}
