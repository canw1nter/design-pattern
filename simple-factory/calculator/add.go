package calculator

type AddOperation struct {
	*Operation
}

func (addOperation *AddOperation) GetResult() float64 {
	return addOperation.num1 + addOperation.num2
}
