package main

import "fmt"

// named return: nama hasil dan err sudah dideklarasikan; bare return mengisi mereka.
func bagiNamed(a, b int) (hasil int, err error) {
	if b == 0 {
		err = fmt.Errorf("pembagi nol")
		return // bare return — hati-hati di fungsi panjang (mudah membingungkan)
	}
	hasil = a / b
	return
}

func main() {
	x, err := bagiNamed(9, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("hasil:", x)
}
