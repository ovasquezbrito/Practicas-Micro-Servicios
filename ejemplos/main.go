package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "¡Guau!"
}

type Cat struct{}

func (c *Cat) Speak() string {
	return "¡Miau!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "??????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Desing patterns!"
}

func main() {
	animals := []Animal{new(Dog), new(Cat), Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
