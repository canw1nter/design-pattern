package cashier

type ICashStrategy interface {
	CashAlgorithm(total float64) float64
}
