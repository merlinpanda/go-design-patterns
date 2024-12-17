package main

import (
	"fmt"
	"github.com/merlinpanda/go-design-pattern/factory"
)

func getGun(gunType string) (factory.IGun, error) {
	if gunType == "ak47" {
		return factory.NewAk47(), nil
	}
	if gunType == "musket" {
		return factory.NewMusket(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
	ak47, err := getGun("ak47")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ak47.GetName())
	fmt.Println(ak47.GetPower())
	ak47.SetName("Ak47 new name")
	ak47.SetPower(12)
	fmt.Println(ak47.GetName())
	fmt.Println(ak47.GetPower())
}
