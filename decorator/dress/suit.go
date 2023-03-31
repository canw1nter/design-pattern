package dress

import "fmt"

type Suit struct {
	*Finery
}

func SuitConstructor(component IPerson) *Suit {
	return &Suit{
		&Finery{component},
	}
}

func (suit *Suit) Show() {
	fmt.Print("西装\t")
	suit.component.Show()
}
