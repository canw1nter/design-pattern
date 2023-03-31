package calculator

type OperationFactory struct{}

// todo singleton factory
var singleOperationFactory = &OperationFactory{}

func OperationFactoryConstructor() *OperationFactory {
	return singleOperationFactory
}

func (operationFactory *OperationFactory) CreateOperation(operator string) IOperation {
	switch operator {
	case "+":
		return &AddOperation{&Operation{}}
	case "-":
		return &SubOperation{&Operation{}}
	default:
		return &Operation{}
	}
}
