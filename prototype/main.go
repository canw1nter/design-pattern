package main

import (
	"design-pattern/prototype/resume"
	"fmt"
)

func main() {
	canwinter := resume.PeopleResumeConstructor("canwinter", "12",
		"company", "time")

	canwinter2 := canwinter.Clone().(resume.PeopleResume)

	canwinter2.Name = "canwinter2"
	canwinter2.Work.Company = "company2"

	fmt.Println(canwinter.Name, canwinter.Age, canwinter.Work)
	fmt.Println(canwinter2.Name, canwinter2.Age, canwinter2.Work)
}
