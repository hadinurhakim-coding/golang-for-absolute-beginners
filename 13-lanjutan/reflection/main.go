package main

import (
	"fmt"
	"reflect"
)

type Orang struct {
	Nama string
	Umur int
}

func main() {
	o := Orang{Nama: "Andi", Umur: 20}
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	fmt.Println("Type:", t.Name(), "Kind:", t.Kind())

	f := v.FieldByName("Nama")
	fmt.Println("Field Nama:", f.String())
}
