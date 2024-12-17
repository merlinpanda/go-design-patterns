package main

import (
	"fmt"
	"github.com/merlinpanda/go-design-pattern/abstract_factory"
)

func main() {
	adidasFactory := &abstract_factory.Adidas{}
	nikeFactory := &abstract_factory.Nike{}

	adidasShoe := adidasFactory.MakeShoe()
	fmt.Println(adidasShoe.GetLogo())
	fmt.Println(adidasShoe.GetSize())

	adidasShirt := adidasFactory.MakeShirt()
	fmt.Println(adidasShirt.GetLogo())
	fmt.Println(adidasShirt.GetSize())

	nikeShoe := nikeFactory.MakeShoe()
	fmt.Println(nikeShoe.GetLogo())
	fmt.Println(nikeShoe.GetSize())

	nikeShirt := nikeFactory.MakeShirt()
	fmt.Println(nikeShirt.GetLogo())
	fmt.Println(nikeShirt.GetSize())
}
