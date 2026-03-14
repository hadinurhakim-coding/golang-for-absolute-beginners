// =============================================================================
// DEKLARASI TIPE (TYPE DECLARATION) DI GO
// =============================================================================
// Type declaration = memberi nama baru untuk tipe (bisa berdasarkan tipe yang sudah ada).
// Bentuk: type NamaTipe TipeDasar  atau  type NamaTipe = TipeYangSama
// Manfaat: kode lebih jelas, type safety, dan bisa lampirkan method ke tipe baru.
// =============================================================================

package main

import "fmt"

func main() {
	fmt.Println("========== 1. Type definition: tipe baru (berbeda dari dasar) ==========")
	typeDefinition()

	fmt.Println("\n========== 2. Type alias: nama lain, tipe sama ==========")
	typeAlias()

	fmt.Println("\n========== 3. Struct: tipe bentukan (beberapa field) ==========")
	structType()

	fmt.Println("\n========== 4. Tipe untuk konstanta / semantik ==========")
	semanticType()

	fmt.Println("\n========== 5. Konversi antara tipe bentukan dan tipe dasar ==========")
	conversionCustomType()
}

// -----------------------------------------------------------------------------
// 1. type Nama TipeDasar → tipe BARU (distinct type)
//    Meski dasarnya int, MyInt dan int tidak bisa dipakai saling menggantikan tanpa konversi.
// -----------------------------------------------------------------------------
func typeDefinition() {
	type Meter int
	type Kilogram int

	var panjang Meter = 10
	var berat Kilogram = 5

	fmt.Printf("panjang = %d Meter, berat = %d Kilogram\n", panjang, berat)
	// panjang = berat  // Error: tipe beda (type safety)
	// var x int = panjang  // Error: perlu konversi eksplisit int(panjang)

	// Konversi ke int bila perlu
	fmt.Printf("panjang sebagai int = %d\n", int(panjang))
}

// -----------------------------------------------------------------------------
// 2. type Nama = TipeYangSama → alias (bukan tipe baru)
//    Nama dan tipe kanan sama persis; bisa dipakai saling menggantikan.
// -----------------------------------------------------------------------------
func typeAlias() {
	type MyInt = int // MyInt adalah alias untuk int, bukan tipe baru

	var a int = 42
	var b MyInt = 42

	// a dan b bisa dipakai saling menggantikan (tipe sama)
	fmt.Printf("int = %d, MyInt = %d\n", a, b)
	a = b   // OK
	b = a   // OK
	fmt.Printf("Setelah a=b, b=a: a=%d, b=%d\n", a, b)
}

// -----------------------------------------------------------------------------
// 3. type Nama struct { Field Tipe ... } → tipe bentukan dengan field
// -----------------------------------------------------------------------------
func structType() {
	type Orang struct {
		Nama string
		Umur int
	}

	p := Orang{Nama: "Budi", Umur: 20}
	fmt.Printf("Orang: %+v\n", p)
	fmt.Printf("Nama: %s, Umur: %d\n", p.Nama, p.Umur)
}

// -----------------------------------------------------------------------------
// 4. Tipe bentukan untuk makna (semantik): ID, Currency, dll.
// -----------------------------------------------------------------------------
func semanticType() {
	type UserID int
	type OrderID int

	var uid UserID = 1001
	var oid OrderID = 5001

	fmt.Printf("UserID = %d, OrderID = %d\n", uid, oid)
	// uid = oid  // Error: UserID dan OrderID beda tipe (mencegah tertukar)
}

// -----------------------------------------------------------------------------
// 5. Konversi antara tipe bentukan (definition) dan tipe dasar
// -----------------------------------------------------------------------------
func conversionCustomType() {
	type Meter int
	m := Meter(100)
	i := int(m) // Meter → int
	fmt.Printf("Meter(100) → int = %d\n", i)
	fmt.Printf("int(50) → Meter = %d\n", Meter(50))
}
