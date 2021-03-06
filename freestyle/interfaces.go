package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

// define interfaces
//type Men interface {
//	SayHi()
//	Sing(lyrics string)
//	Guzzle(beerStein string)
//}
//
//type YoungChap interface {
//	SayHi()
//	Sing(song string)
//	BorrowMoney(amount float32)
//}
//
//type ElderlyGent interface {
//	SayHi()
//	Sing(song string)
//	SpendSalary(amount float32)
//}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee overloads SayHi
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

func main() {
	var stud Student = Student{Human{"Alexander", 22, "+7(951)151-05-51"}, "BSTU", 4000}
	var em Employee = Employee{Human{"Alex", 31, "+7(977)851-97-99"}, "ExactFarming", 140000}

	stud.SayHi()
	stud.BorrowMoney(1000)

	em.SayHi()
	em.SpendSalary(25000)
}
