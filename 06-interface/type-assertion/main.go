package main

import "fmt"

type Produk struct {
	Nama  string
	Harga int
}

func main() {
	var x any = "pesan teks"

	// Type assertion satu nilai: x.(T). Hati-hati — gagal tipe → panic saat runtime.
	s := x.(string)
	fmt.Println("string:", s)

	// Comma-ok: aman; ok false jika nilai di x bukan bertipe T.
	x = 42
	if n, ok := x.(int); ok {
		fmt.Println("int:", n)
	} else {
		fmt.Println("bukan int")
	}

	x = "bukan angka"
	if _, ok := x.(int); !ok {
		fmt.Println("x saat ini bukan int (comma-ok menghindari panic)")
	}

	// switch pada tipe: pola umum untuk banyak kemungkinan tipe.
	tunjukkanTipe("hello")
	tunjukkanTipe(7)
	tunjukkanTipe(Produk{Nama: "buku", Harga: 50000})
	tunjukkanTipe(3.14)
}

// tunjukkanTipe membedakan perilaku berdasarkan tipe konkret di balik any.
func tunjukkanTipe(x any) {
	switch v := x.(type) {
	case string:
		fmt.Println("string, panjang:", len(v))
	case int:
		fmt.Println("int, nilai:", v)
	case Produk:
		fmt.Println("Produk:", v.Nama, v.Harga)
	default:
		fmt.Printf("tipe lain: %T\n", v)
	}
}
