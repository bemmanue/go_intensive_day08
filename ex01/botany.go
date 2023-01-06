package main

import (
	"fmt"
	"reflect"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(plant interface{}) {
	pType := reflect.TypeOf(plant)
	pVal := reflect.ValueOf(plant)

	for i := 0; i < pType.NumField(); i++ {
		field := pType.Field(i)
		fmt.Print(field.Name, "")
		if len(field.Tag) != 0 {
			fmt.Print("(", field.Tag, ")")
		}
		fmt.Println(":", pVal.FieldByName(field.Name))
	}
}

func main() {
	var plant = UnknownPlant{"maple", "hand shaped", 231}
	var plant2 = AnotherUnknownPlant{10, "lanceolate", 15}

	describePlant(plant)
	fmt.Println()
	describePlant(plant2)
}
