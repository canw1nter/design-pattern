package factory

import "design-pattern/factory-method/calculator"

type SubOperationFactory struct {
}

func SubOperationFactoryConstructor() *SubOperationFactory {
	return &SubOperationFactory{}
}

func (*SubOperationFactory) CreateOperation() calculator.IOperation {
	return &calculator.SubOperation{Operation: &calculator.Operation{}}
}
