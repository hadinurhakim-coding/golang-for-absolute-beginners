package main

import (
	"encoding/json"
	"fmt"
)

type Orang struct {
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
}

func main() {
	o := Orang{Nama: "Andi", Umur: 20}
	b, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	raw := `{"nama":"Budi","umur":25}`
	var p Orang
	if err := json.Unmarshal([]byte(raw), &p); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)
}
