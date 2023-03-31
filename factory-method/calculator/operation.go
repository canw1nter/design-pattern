package calculator

type IOperation interface {
	// getter

	GetNum1() float64
	GetNum2() float64

	// setter

	SetNum1(num float64)
	SetNum2(num float64)

	GetResult() float64
}

type Operation struct {
	num1 float64
	num2 float64
}

func (operation *Operation) GetNum1() float64 {
	return operation.num1
}

func (operation *Operation) GetNum2() float64 {
	return operation.num2
}

func (operation *Operation) SetNum1(num float64) {
	operation.num1 = num
}

func (operation *Operation) SetNum2(num float64) {
	operation.num2 = num
}

func (operation *Operation) GetResult() float64 {
	return 0
}
