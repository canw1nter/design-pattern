package main

import "design-pattern/proxy/pursue"

func main() {
	pursued := pursue.PursuedConstructor("被追求者")
	proxy := pursue.ProxyConstructor("追求者", pursued)

	proxy.GiveDolls()
	proxy.GiveFlowers()
	proxy.GiveChocolate()
}
