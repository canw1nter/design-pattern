package cashier

type RebateCashStrategy struct {
	discount float64
}

func (rebate *RebateCashStrategy) CashAlgorithm(total float64) float64 {
	return total * rebate.discount
}
