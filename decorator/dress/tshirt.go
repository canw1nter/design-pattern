package dress

import "fmt"

type TShirt struct {
	*Finery
}

func TShirtConstructor(component IPerson) *TShirt {
	return &TShirt{
		&Finery{component},
	}
}

func (tShirt *TShirt) Show() {
	fmt.Print("T恤\t")
	tShirt.component.Show()
}
