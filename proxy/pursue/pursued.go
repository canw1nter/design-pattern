package pursue

type Pursued struct {
	name string
}

func PursuedConstructor(name string) *Pursued {
	return &Pursued{
		name: name,
	}
}
