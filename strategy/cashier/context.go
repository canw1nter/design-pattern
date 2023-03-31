package cashier

type CashContext struct {
	strategy ICashStrategy
}

func CashContextConstructor(strategyType int) *CashContext {
	context := &CashContext{}

	switch strategyType {
	case 1:
		context.strategy = &NormalCashStrategy{}
	case 2:
		// some code for acquire discount
		context.strategy = &RebateCashStrategy{discount: 0.8}
	case 3:
		// some code for acquire condition and back
		context.strategy = &CashbackCashStrategy{condition: 400, back: 50}
	default:
		context.strategy = &NormalCashStrategy{}
	}

	return context
}

func (cashContext *CashContext) GetFinalAmount(total float64) float64 {
	return cashContext.strategy.CashAlgorithm(total)
}
