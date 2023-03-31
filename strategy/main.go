package main

import (
	"design-pattern/strategy/cashier"
	"fmt"
)

func main() {
	fmt.Println("选择优惠方式：1.正常收费；2.打折；3.返现")

	var strategyType int
	fmt.Scanln(&strategyType)

	var price float64
	var count int
	fmt.Println("输入单价和数量")
	fmt.Scanln(&price, &count)

	context := cashier.CashContextConstructor(strategyType)

	fmt.Println("应付金额：", context.GetFinalAmount(price*float64(count)))
}
