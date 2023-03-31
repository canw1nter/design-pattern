package dress

type Finery struct {
	component IPerson
}

func (finery *Finery) Show() {
	finery.component.Show()
}
