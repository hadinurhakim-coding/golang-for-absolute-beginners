// =============================================================================
// OPERASI MATEMATIKA DI GO
// =============================================================================
// Operator: + (tambah), - (kurang), * (kali), / (bagi), % (modulo/sisa bagi)
// Integer division: 7/2 = 3 (hasil bagi bulat). Untuk float: 7.0/2.0 atau float64(7)/2
// =============================================================================

package main

import "fmt"

func main() {
	fmt.Println("========== 1. Dasar: + - * / % ==========")
	dasar()

	fmt.Println("\n========== 2. Pembagian integer vs float ==========")
	pembagian()

	fmt.Println("\n========== 3. Modulo (sisa bagi) ==========")
	modulo()

	fmt.Println("\n========== 4. Compound assignment: +=, -=, *=, /=, %= ==========")
	compoundAssignment()

	fmt.Println("\n========== 5. Increment & decrement (++, --) ==========")
	incrementDecrement()

	fmt.Println("\n========== 6. Urutan operasi (precedence) ==========")
	precedence()

	fmt.Println("\n========== 7. Contoh kasus: total belanja & diskon ==========")
	kasusBelanja()

	fmt.Println("\n========== 8. Contoh kasus: konversi suhu ==========")
	kasusSuhu()
}

func dasar() {
	a, b := 10, 3
	fmt.Printf("a=%d, b=%d\n", a, b)
	fmt.Printf("a + b = %d\n", a+b)
	fmt.Printf("a - b = %d\n", a-b)
	fmt.Printf("a * b = %d\n", a*b)
	fmt.Printf("a / b = %d (pembagian bulat)\n", a/b)
	fmt.Printf("a %% b = %d (sisa bagi)\n", a%b)
}

func pembagian() {
	fmt.Println("7 / 2 (int)     =", 7/2)
	fmt.Println("7.0 / 2 (float)  =", 7.0/2)
	fmt.Println("float64(7) / 2  =", float64(7)/2)
}

func modulo() {
	fmt.Println("10 % 3 =", 10%3)
	fmt.Println("15 % 5 =", 15%5)
	fmt.Println("7 % 2  =", 7%2, "(cek ganjil: sisa 1 = ganjil)")
}

func compoundAssignment() {
	n := 10
	n += 5
	fmt.Printf("n = 10, n += 5 → n = %d\n", n)
	n -= 3
	fmt.Printf("n -= 3 → n = %d\n", n)
	n *= 2
	fmt.Printf("n *= 2 → n = %d\n", n)
	n /= 4
	fmt.Printf("n /= 4 → n = %d\n", n)
	n %= 3
	fmt.Printf("n %%= 3 → n = %d\n", n)
}

func incrementDecrement() {
	x := 5
	x++
	fmt.Printf("x = 5, x++ → x = %d\n", x)
	x--
	fmt.Printf("x-- → x = %d\n", x)
	// Di Go: ++ dan -- hanya postfix (x++), tidak ada prefix (++x). Tidak bisa dipakai dalam ekspresi: z := x++ // error
}

func precedence() {
	// * dan / lebih dulu dari + dan -
	fmt.Println("2 + 3*4   =", 2+3*4)
	fmt.Println("(2 + 3)*4 =", (2+3)*4)
}

func kasusBelanja() {
	hargaItem1 := 25000
	hargaItem2 := 15000
	jumlahItem1 := 2
	jumlahItem2 := 3

	subtotal := hargaItem1*jumlahItem1 + hargaItem2*jumlahItem2
	diskonPersen := 10
	diskon := subtotal * diskonPersen / 100
	total := subtotal - diskon

	fmt.Printf("Item1: %d x %d = %d\n", hargaItem1, jumlahItem1, hargaItem1*jumlahItem1)
	fmt.Printf("Item2: %d x %d = %d\n", hargaItem2, jumlahItem2, hargaItem2*jumlahItem2)
	fmt.Printf("Subtotal: %d\n", subtotal)
	fmt.Printf("Diskon %d%%: %d\n", diskonPersen, diskon)
	fmt.Printf("Total bayar: %d\n", total)
}

func kasusSuhu() {
	celsius := 25.0
	fahrenheit := celsius*9/5 + 32
	fmt.Printf("%.1f°C = %.1f°F\n", celsius, fahrenheit)

	km := 100.0
	mil := km / 1.609
	fmt.Printf("%.1f km = %.2f mil\n", km, mil)
}
