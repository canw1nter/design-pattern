package main

import (
	"design-pattern/simple-factory/calculator"
	"fmt"
)

func main() {
	var operator string
	fmt.Println("请输入运算符")
	fmt.Scanln(&operator)

	var num1, num2 float64
	fmt.Println("请输入两个数字")
	fmt.Scan(&num1, &num2)

	resultOperation := calculator.OperationFactoryConstructor().CreateOperation(operator)

	resultOperation.SetNum1(num1)
	resultOperation.SetNum2(num2)

	fmt.Println("运算结果为")
	fmt.Println(resultOperation.GetResult())
}
