// =============================================================================
// TIPE DATA BOOLEAN DI GO
// =============================================================================
// Di Go hanya ada satu tipe untuk nilai benar/salah: bool.
// Nilai yang mungkin: true atau false (bukan 1/0 seperti di C).
// Dipakai untuk kondisi (if, for), operator logika (&&, ||, !), dan perbandingan.
// =============================================================================

package main

import "fmt"

func main() {
	fmt.Println("========== 1. Dasar: true & false ==========")
	dasar()

	fmt.Println("\n========== 2. Zero value bool ==========")
	zeroValue()

	fmt.Println("\n========== 3. Operator logika: &&, ||, ! ==========")
	operatorLogika()

	fmt.Println("\n========== 4. Hasil perbandingan (==, !=, <, >, <=, >=) ==========")
	perbandingan()

	fmt.Println("\n========== 5. Short-circuit (evaluasi singkat) ==========")
	shortCircuit()

	fmt.Println("\n========== 6. Bool dalam kondisi if ==========")
	kondisiIf()
}

// -----------------------------------------------------------------------------
// 1. Deklarasi dan nilai literal: true, false
// -----------------------------------------------------------------------------
func dasar() {
	var ya bool = true
	var tidak bool = false

	fmt.Printf("ya   = %t (tipe: bool)\n", ya)
	fmt.Printf("tidak= %t (tipe: bool)\n", tidak)

	// Inferensi tipe
	benar := true
	salah := false
	fmt.Printf("benar= %t, salah= %t\n", benar, salah)
}

// -----------------------------------------------------------------------------
// 2. Zero value: variabel bool yang belum diisi = false
// -----------------------------------------------------------------------------
func zeroValue() {
	var b bool
	fmt.Printf("var b bool → zero value = %t\n", b)
}

// -----------------------------------------------------------------------------
// 3. Operator logika
//   && = AND (dan)  : true hanya jika keduanya true
//   || = OR (atau)  : true jika salah satu true
//   !  = NOT (negasi): membalik true↔false
// -----------------------------------------------------------------------------
func operatorLogika() {
	fmt.Println("AND (&&): true && true =", true && true)   // true
	fmt.Println("AND (&&): true && false=", true && false)   // false
	fmt.Println("AND (&&): false && false=", false && false) // false

	fmt.Println("OR (||) : true || false =", true || false)  // true
	fmt.Println("OR (||) : false || false=", false || false) // false

	fmt.Println("NOT (!) : !true =", !true)   // false
	fmt.Println("NOT (!) : !false=", !false)  // true
}

// -----------------------------------------------------------------------------
// 4. Hasil perbandingan selalu bool
//   ==  sama dengan
//   !=  tidak sama dengan
//   <   kurang dari
//   >   lebih dari
//   <=  kurang dari atau sama dengan
//   >=  lebih dari atau sama dengan
// -----------------------------------------------------------------------------
func perbandingan() {
	a, b := 10, 20

	fmt.Printf("a=%d, b=%d\n", a, b)
	fmt.Printf("a == b  → %t\n", a == b)  // false
	fmt.Printf("a != b  → %t\n", a != b)  // true
	fmt.Printf("a < b   → %t\n", a < b)   // true
	fmt.Printf("a > b   → %t\n", a > b)   // false
	fmt.Printf("a <= 10 → %t\n", a <= 10) // true
	fmt.Printf("a >= 10 → %t\n", a >= 10) // true

	// String juga bisa dibandingkan (urutan leksikografis)
	fmt.Printf("\n\"abc\" == \"abc\" → %t\n", "abc" == "abc")
	fmt.Printf("\"abc\" < \"abd\"  → %t\n", "abc" < "abd")
}

// -----------------------------------------------------------------------------
// 5. Short-circuit: && berhenti di false pertama; || berhenti di true pertama
// -----------------------------------------------------------------------------
func shortCircuit() {
	// Fungsi yang mengembalikan bool dan "efek samping" (di sini kita cek urutan pemanggilan)
	call := func(name string, result bool) bool {
		fmt.Printf("  [dipanggil: %s]\n", name)
		return result
	}

	fmt.Println("false && call(\"B\", true) → B tidak perlu dievaluasi:")
	_ = false && call("B", true)

	fmt.Println("true || call(\"Y\", false) → Y tidak perlu dievaluasi:")
	_ = true || call("Y", false)
}

// -----------------------------------------------------------------------------
// 6. Bool dipakai langsung di if (tidak pakai perbandingan dengan 0/1)
// -----------------------------------------------------------------------------
func kondisiIf() {
	login := true
	if login {
		fmt.Println("  if login { ... } → kondisi true, blok ini dijalankan.")
	}

	admin := false
	if !admin {
		fmt.Println("  if !admin { ... } → admin false, !admin true, blok dijalankan.")
	}
}
