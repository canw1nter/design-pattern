package dress

import "fmt"

type IPerson interface {
	Show()
}

type Person struct {
	name string
}

func PersonConstructor(name string) *Person {
	return &Person{name}
}

func (person *Person) Show() {
	fmt.Print("装饰的：", person.name)
}
