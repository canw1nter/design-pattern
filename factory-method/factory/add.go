package factory

import "design-pattern/factory-method/calculator"

type AddOperationFactory struct {
}

func AddOperationFactoryConstructor() *AddOperationFactory {
	return &AddOperationFactory{}
}

func (*AddOperationFactory) CreateOperation() calculator.IOperation {
	return &calculator.AddOperation{Operation: &calculator.Operation{}}
}
