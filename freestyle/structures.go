package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	firstName string
	lastName string
	gender string
	phone int
}

func (p *Person) setName(name string) {
	p.firstName = name
}

func (p *Person) setPhone(phone string) {
	result := strings.Replace(phone, "+", "", -1)
	result = strings.Replace(result, "(", "", -1)
	result = strings.Replace(result, ")", "", -1)
	result = strings.Replace(result, "-", "", -1)
	result = strings.Replace(result, " ", "", -1)
	resultInt, err := strconv.Atoi(result)
	if err == nil {
		p.phone = resultInt
	} else {
		fmt.Println(err)
	}
}

type Account struct {
	email string
	password string
}

type User struct {
	Person
	Account
}

func main() {
	var user1 User = User{
		Person{"Alexander", "Mospan", "M", 0},
		Account{"alexander.mospan@gmail.com", "qwe123"},
	}

	//fmt.Println(user1.firstName)
	//user1.setName("Alex")
	//fmt.Println(user1.firstName)

	user1.setPhone("+7(977)851-97 99")

	fmt.Println(user1.phone) // 79778519799
}