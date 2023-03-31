package calculator

type SubOperation struct {
	*Operation
}

func (subOperation *SubOperation) GetResult() float64 {
	return subOperation.num2 - subOperation.num1
}
