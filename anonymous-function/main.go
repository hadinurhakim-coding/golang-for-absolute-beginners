// =============================================================================
// ANONYMOUS FUNCTION DI GO
// =============================================================================
// Anonymous function = fungsi tanpa nama (literal function). Bentuk: func(param T) R { }
// atau func() { }. Bisa di-assign ke variabel, dikirim sebagai parameter, atau
// dipanggil langsung (IIFE). Berguna untuk logic sekali pakai atau aturan yang
// berubah-ubah tanpa perlu fungsi permanen.
//
// POLA KOMENTAR (konsisten):
//   1. Penjelasan        — ringkasan materi / apa yang dilakukan.
//   2. Alasan penggunaan — kenapa fitur ini dipakai.
//   3. Fungsi penggunaan — cara pakai (assign, pass, IIFE).
//   4. Output            — contoh hasil cetakan/return dari fungsi tersebut.
//   5. Penjelasan output — mengapa output-nya bisa seperti itu (alur/logika).
// =============================================================================

package main

import "fmt"

func main() {
	fmt.Println("========== 1. Dasar: anonymous function ==========")
	dasarAnonymous()

	fmt.Println("\n========== 2. Anonymous function sebagai parameter ==========")
	anonymousSebagaiParameter()

	fmt.Println("\n========== 3. Studi kasus: total belanja + diskon fleksibel ==========")
	kasusTotalBelanjaDiskon()

	fmt.Println("\n========== 4. Rincian: pelanggan dengan jenis diskon berbeda ==========")
	rincianPelangganDiskon()
}

// -----------------------------------------------------------------------------
// 1. Dasar: anonymous function
//
// Penjelasan:
//   Fungsi tanpa nama didefinisikan inline. Bisa disimpan di variabel:
//   inc := func(x int) int { return x + 1 }. Atau dipanggil langsung:
//   func() { fmt.Println("hi") }() — tanda () di akhir menjalankan fungsi.
//
// Alasan penggunaan:
//   Logic yang hanya dipakai sekali atau di satu tempat; tidak perlu
//   mendefinisikan fungsi bernama di level package.
//
// Fungsi penggunaan:
//   inc := func(x int) int { return x + 1 }; inc(5) = 6.
//
// Output:
//   inc(5) = 6
//   IIFE: Hallo
//
// Penjelasan output:
//   inc adalah literal yang return x+1; inc(5) return 6. IIFE func(){ ... }()
//   langsung dijalankan sehingga mencetak "Hallo".
// -----------------------------------------------------------------------------
func dasarAnonymous() {
	inc := func(x int) int { return x + 1 }
	fmt.Printf("  inc(5) = %d\n", inc(5))
	fmt.Print("  IIFE: ")
	func() { fmt.Println("Hallo") }()
}

// -----------------------------------------------------------------------------
// 2. Anonymous function sebagai parameter
//
// Penjelasan:
//   Parameter yang bertipe fungsi bisa diisi dengan literal (anonymous).
//   Pemanggil menulis func(...) { ... } langsung di dalam argument, tanpa
//   mendefinisikan fungsi bernama terlebih dahulu.
//
// Alasan penggunaan:
//   Aturan atau perilaku yang spesifik untuk satu pemanggilan; tidak perlu
//   mengotori namespace dengan banyak fungsi kecil.
//
// Fungsi penggunaan:
//   apply(10, func(x int) int { return x * 3 }) = 30. Literal jadi parameter.
//
// Output:
//   apply(10, func(x int) int { return x*3 }) = 30
//
// Penjelasan output:
//   apply memanggil fn(10) dengan fn = fungsi literal return x*3; 10*3=30,
//   sehingga output "30".
// -----------------------------------------------------------------------------
func anonymousSebagaiParameter() {
	result := apply(10, func(x int) int { return x * 3 })
	fmt.Printf("  apply(10, func(x int) int { return x*3 }) = %d\n", result)
}

func apply(n int, fn func(int) int) int {
	return fn(n)
}

// -----------------------------------------------------------------------------
// 3. Studi kasus: total belanja + diskon fleksibel
//
// Penjelasan:
//   HitungTotalBelanja(hargas []int, diskonFn func(subtotal int) int) menjumlahkan
//   harga lalu menerapkan aturan diskon lewat parameter diskonFn. Setiap jenis
//   diskon (member, cuci gudang, tanpa diskon) diwakili oleh anonymous function
//   yang berbeda — tidak perlu fungsi permanen per jenis diskon.
//
// Alasan penggunaan:
//   Aturan diskon beragam dan bisa berubah (promo baru, kebijakan toko).
//   Anonymous function memungkinkan satu fungsi HitungTotalBelanja dipakai
//   dengan banyak aturan tanpa menambah fungsi bernama (diskonMember, diskonGudang, dll).
//
// Fungsi penggunaan:
//   Tanpa diskon: diskonFn = func(s int) int { return s }
//   Member 10%:  diskonFn = func(s int) int { return s * 90 / 100 }
//   Cuci gudang 25%: diskonFn = func(s int) int { return s * 75 / 100 }
//
// Output:
//   Subtotal: 100000
//   Tanpa diskon: 100000
//   Diskon member 10%: 90000
//   Diskon cuci gudang 25%: 75000
//
// Penjelasan output:
//   Subtotal 100000. Tanpa diskon: diskonFn return 100000. Member 10%: return
//   100000*90/100=90000. Cuci gudang 25%: return 100000*75/100=75000.
// -----------------------------------------------------------------------------
func kasusTotalBelanjaDiskon() {
	hargas := []int{25000, 35000, 40000} // 100.000
	subtotal := sumSlice(hargas)
	fmt.Printf("  Subtotal: %d\n", subtotal)

	// Tanpa diskon: anonymous function yang mengembalikan nilai asli
	tanpaDiskon := func(s int) int { return s }
	total1 := HitungTotalBelanja(hargas, tanpaDiskon)
	fmt.Printf("  Tanpa diskon: %d\n", total1)

	// Diskon member 10%
	diskonMember10 := func(s int) int { return s * 90 / 100 }
	total2 := HitungTotalBelanja(hargas, diskonMember10)
	fmt.Printf("  Diskon member 10%%: %d\n", total2)

	// Diskon cuci gudang 25%
	diskonCuciGudang25 := func(s int) int { return s * 75 / 100 }
	total3 := HitungTotalBelanja(hargas, diskonCuciGudang25)
	fmt.Printf("  Diskon cuci gudang 25%%: %d\n", total3)
}

func sumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

// HitungTotalBelanja menjumlahkan harga lalu menerapkan aturan diskon via diskonFn.
// diskonFn(subtotal) mengembalikan jumlah yang harus dibayar setelah diskon.
func HitungTotalBelanja(hargas []int, diskonFn func(subtotal int) int) int {
	subtotal := sumSlice(hargas)
	return diskonFn(subtotal)
}

// -----------------------------------------------------------------------------
// 4. Rincian: pelanggan dengan jenis diskon berbeda
//
// Penjelasan:
//   Tiga pelanggan: A (tanpa diskon), B (member 10%), C (cuci gudang 25%).
//   Masing-masing memakai anonymous function berbeda sebagai parameter diskon.
//   Satu fungsi HitungTotalBelanja dipakai untuk ketiganya; aturan diskon
//   ditentukan di tempat pemanggilan tanpa fungsi permanen.
//
// Alasan penggunaan:
//   Menunjukkan fleksibilitas: setiap pelanggan bisa punya aturan diskon
//   berbeda (dari database, config, atau pilihan UI) dengan mengirim
//   anonymous function yang sesuai.
//
// Fungsi penggunaan:
//   HitungTotalBelanja(itemsA, func(s int) int { return s })
//   HitungTotalBelanja(itemsB, func(s int) int { return s * 90 / 100 })
//   HitungTotalBelanja(itemsC, func(s int) int { return s * 75 / 100 })
//
// Output:
//   Pelanggan A (tanpa diskon): 150000 -> 150000
//   Pelanggan B (member 10%): 150000 -> 135000
//   Pelanggan C (cuci gudang 25%): 150000 -> 112500
//
// Penjelasan output:
//   A: subtotal 150000, diskonFn return s = 150000. B: 150000*90/100=135000.
//   C: 150000*75/100=112500. Tiap baris memakai anonymous function berbeda.
// -----------------------------------------------------------------------------
func rincianPelangganDiskon() {
	itemsA := []int{50000, 50000, 50000} // 150.000
	itemsB := []int{50000, 50000, 50000}
	itemsC := []int{50000, 50000, 50000}

	// Pelanggan A: tanpa diskon (anonymous function inline)
	totalA := HitungTotalBelanja(itemsA, func(s int) int { return s })
	fmt.Printf("  Pelanggan A (tanpa diskon): 150000 -> %d\n", totalA)

	// Pelanggan B: diskon member 10% (anonymous function inline)
	totalB := HitungTotalBelanja(itemsB, func(s int) int { return s * 90 / 100 })
	fmt.Printf("  Pelanggan B (member 10%%): 150000 -> %d\n", totalB)

	// Pelanggan C: diskon cuci gudang 25% (anonymous function inline)
	totalC := HitungTotalBelanja(itemsC, func(s int) int { return s * 75 / 100 })
	fmt.Printf("  Pelanggan C (cuci gudang 25%%): 150000 -> %d\n", totalC)
}
