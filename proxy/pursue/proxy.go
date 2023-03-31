package pursue

import "fmt"

type Proxy struct {
	pursuing *Pursuing
}

func ProxyConstructor(name string, pursued *Pursued) *Proxy {
	return &Proxy{
		pursuing: &Pursuing{name: name, pursued: pursued},
	}
}

func (proxy *Proxy) GiveDolls() {
	fmt.Println("代理告诉", proxy.pursuing.name, proxy.pursuing.pursued.name, "想要洋娃娃")
	proxy.pursuing.GiveDolls()
	fmt.Println("代理把", proxy.pursuing.name, "的洋娃娃送给", proxy.pursuing.pursued.name)
}

func (proxy *Proxy) GiveFlowers() {
	fmt.Println("代理告诉", proxy.pursuing.name, proxy.pursuing.pursued.name, "想要花")
	proxy.pursuing.GiveFlowers()
	fmt.Println("代理把", proxy.pursuing.name, "的花送给", proxy.pursuing.pursued.name)
}

func (proxy *Proxy) GiveChocolate() {
	fmt.Println("代理告诉", proxy.pursuing.name, proxy.pursuing.pursued.name, "想要巧克力")
	proxy.pursuing.GiveChocolate()
	fmt.Println("代理把", proxy.pursuing.name, "的巧克力送给", proxy.pursuing.pursued.name)
}
