package resume

type WorkExperience struct {
	Company  string
	TimeArea string
}

type PeopleResume struct {
	Name string
	Age  string
	Work *WorkExperience
}

func WorkExperienceConstructor(company, timeArea string) *WorkExperience {
	return &WorkExperience{
		company,
		timeArea,
	}
}

func PeopleResumeConstructor(name, age, company, timeArea string) PeopleResume {
	return PeopleResume{
		name,
		age,
		WorkExperienceConstructor(company, timeArea),
	}
}

func (resume PeopleResume) Clone() IPrototype {
	return PeopleResumeConstructor(resume.Name, resume.Age, resume.Work.Company, resume.Work.TimeArea)
}
