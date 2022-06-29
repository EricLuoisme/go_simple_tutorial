package main

import (
	"fmt"
	"reflect"
)

// Doctor upper-case, can be accessed from other pkg
type Doctor struct {
	// can be accessed from other pkg
	Id int
	// can not be accessed from other pkg
	name       string
	companions []string
}

type Animal struct {
	// placing tag by Reflect pkg
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	// embedded Animal struct into the Bird struct
	Animal
	SpeedKPH float64
	CanFly   bool
}

func main() {
	// better for fill struct, for maintain
	aDoctor := Doctor{
		Id:         15,
		name:       "Tom",
		companions: []string{"abc", "def"}}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions[1])
	// unlike map or slices, 'struct' would copy the entire value
	bDoctor := aDoctor
	bDoctor.name = "Dog"
	fmt.Println(aDoctor)
	fmt.Println(bDoctor)

	// inline struct declaring
	cDoctor := struct{ name string }{name: "John Perter"}
	fmt.Println(cDoctor)

	// struct embedded
	aBird := Bird{}
	aBird.Name = "Emu"
	aBird.Origin = "Australia"
	aBird.SpeedKPH = 46.4
	aBird.CanFly = false
	fmt.Println(aBird)

	cBird := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 46.4,
		CanFly:   false}
	fmt.Println(cBird)

	// Tag, could be used for other field-requirement-validation pkg
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
