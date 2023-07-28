package main

import "fmt"

type person struct {
	name    string
	age     int
	address string
}

func createPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	p.address = "HCM"
	return &p
}
func main() {
	fmt.Println(person{"Bob", 20, "HCM"})

	fmt.Println(person{name: "Alice", age: 18})

	fmt.Println(person{name: "Feeding"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(createPerson("Tran TRanh"))

	s := person{name: "Pop", age: 56, address: "Japan"}
	fmt.Println(s.address)

	sp := &s // create poinnter sp refer s
	sp.name = "Osaka"
	fmt.Println(sp.name)

	dog := struct {
		name   string
		origin string
	}{"Pulldo", "Indo"}

	fmt.Println(dog)

}
