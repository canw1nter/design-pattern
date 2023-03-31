package pursue

import "fmt"

type Pursuing struct {
	name    string
	pursued *Pursued
}

func (pursuing *Pursuing) GiveDolls() {
	fmt.Println(pursuing.name, "为", pursuing.pursued.name, "买了洋娃娃")
}

func (pursuing *Pursuing) GiveFlowers() {
	fmt.Println(pursuing.name, "为", pursuing.pursued.name, "买了花")
}

func (pursuing *Pursuing) GiveChocolate() {
	fmt.Println(pursuing.name, "为", pursuing.pursued.name, "买了巧克力")
}
