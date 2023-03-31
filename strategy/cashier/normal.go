package cashier

type NormalCashStrategy struct {
}

func (normal *NormalCashStrategy) CashAlgorithm(total float64) float64 {
	return total
}
