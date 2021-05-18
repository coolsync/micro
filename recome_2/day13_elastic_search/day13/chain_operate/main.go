package main

import "fmt"

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Worked bool   `json:"worked"`
}

func (s *Person) run() *Person {
	fmt.Printf("%s run\n", s.Name)
	return s
}

func (s *Person) jump() *Person {
	fmt.Printf("%s jump\n", s.Name)
	return s
}

func main() {
	p1 := &Person{
		Name: "hello",
		Age: 12,
		Worked: true,
	}
	p1.run().jump()
}