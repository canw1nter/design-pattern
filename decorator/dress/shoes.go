package dress

import "fmt"

type Shoes struct {
	*Finery
}

func ShoesConstructor(component IPerson) *Shoes {
	return &Shoes{
		&Finery{component},
	}
}

func (shoes *Shoes) Show() {
	fmt.Print("鞋子\t")
	shoes.component.Show()
}
