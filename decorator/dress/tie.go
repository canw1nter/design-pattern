package dress

import "fmt"

type Tie struct {
	*Finery
}

func TieConstructor(component IPerson) *Tie {
	return &Tie{
		&Finery{component},
	}
}

func (tie *Tie) Show() {
	fmt.Print("领带\t")
	tie.component.Show()
}
