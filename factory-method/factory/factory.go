package factory

import "design-pattern/factory-method/calculator"

type OperationFactory interface {
	CreateOperation() calculator.IOperation
}
