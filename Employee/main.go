package main

type Employee interface {
	Language() string
	Age() int
}

type Engineer struct {
	Name     string
	AgeYears int
}

func (e Engineer) Age() int {
	return e.AgeYears
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

func main() {
	var programmers []Employee
	yasin := Engineer{Name: "yasin", AgeYears: 24}
	programmers = append(programmers, yasin)
}
