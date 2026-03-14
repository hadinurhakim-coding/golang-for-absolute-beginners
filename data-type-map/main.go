// =============================================================================
// TIPE DATA MAP DI GO
// =============================================================================
// Map = struktur key–value. Key unik; tiap key memetakan ke satu value.
// Deklarasi: map[KeyType]ValueType. Zero value = nil (belum diinisialisasi).
// Pakai make(map[K]V) atau literal map[K]V{k: v, ...} untuk inisialisasi.
// =============================================================================

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("========== 1. Dasar: deklarasi, make(), literal ==========")
	mapDasar()

	fmt.Println("\n========== 2. Tambah, baca, ubah, hapus (CRUD) ==========")
	mapCRUD()

	fmt.Println("\n========== 3. Cek keberadaan key (ok idiom) ==========")
	mapOkIdiom()

	fmt.Println("\n========== 4. len() dan iterasi range ==========")
	mapLenRange()

	fmt.Println("\n========== 5. Map adalah reference type ==========")
	mapReference()

	fmt.Println("\n========== 6. Studi kasus: buku telepon ==========")
	kasusBukuTelepon()

	fmt.Println("\n========== 7. Studi kasus: hitung kata (word count) ==========")
	kasusWordCount()

	fmt.Println("\n========== 8. Studi kasus: konfigurasi/setting ==========")
	kasusKonfigurasi()

	fmt.Println("\n========== 9. Studi kasus: cache lookup ==========")
	kasusCacheLookup()
}

// -----------------------------------------------------------------------------
// 1. make(map[K]V), literal map[K]V{k1: v1, k2: v2}. Key harus comparable (bukan slice/map).
// -----------------------------------------------------------------------------
func mapDasar() {
	// Inisialisasi dengan make
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println("make + assign:", m)

	// Literal
	umur := map[string]int{
		"Budi":  20,
		"Siti":  22,
		"Agus":  19,
	}
	fmt.Println("literal map:", umur)
}

// -----------------------------------------------------------------------------
// 2. Tambah/ubah: m[k] = v. Baca: m[k]. Hapus: delete(m, k).
// -----------------------------------------------------------------------------
func mapCRUD() {
	m := map[string]int{"x": 10, "y": 20}
	fmt.Println("Awal:", m)

	m["z"] = 30
	fmt.Println("Setelah m[\"z\"]=30:", m)

	m["x"] = 100
	fmt.Println("Setelah m[\"x\"]=100:", m)

	fmt.Println("Baca m[\"y\"]:", m["y"])

	delete(m, "y")
	fmt.Println("Setelah delete(m, \"y\"):", m)
}

// -----------------------------------------------------------------------------
// 3. v, ok := m[key]. ok = true jika key ada, false jika tidak. Aman dari zero value.
// -----------------------------------------------------------------------------
func mapOkIdiom() {
	m := map[string]int{"a": 1, "b": 0}
	v, ok := m["a"]
	fmt.Printf("m[\"a\"] → value=%d, ok=%t\n", v, ok)
	v, ok = m["b"]
	fmt.Printf("m[\"b\"] → value=%d, ok=%t (value 0, tapi key ada)\n", v, ok)
	v, ok = m["c"]
	fmt.Printf("m[\"c\"] → value=%d, ok=%t (key tidak ada, value = zero value int)\n", v, ok)
}

// -----------------------------------------------------------------------------
// 4. len(m) = jumlah key. for k, v := range m { } iterasi key-value (urutan tidak tetap).
// -----------------------------------------------------------------------------
func mapLenRange() {
	m := map[string]int{"satu": 1, "dua": 2, "tiga": 3}
	fmt.Println("len(m) =", len(m))
	fmt.Print("range: ")
	for k, v := range m {
		fmt.Printf("%s:%d ", k, v)
	}
	fmt.Println()
}

// -----------------------------------------------------------------------------
// 5. Map diteruskan by reference; ubah di fungsi mempengaruhi map asli.
// -----------------------------------------------------------------------------
func mapReference() {
	m := map[string]int{"a": 1}
	ubah := func(m map[string]int) {
		m["a"] = 99
		m["b"] = 2
	}
	ubah(m)
	fmt.Println("Setelah panggil ubah(m):", m)
}

// -----------------------------------------------------------------------------
// 6. Studi kasus: buku telepon (nama → nomor)
// -----------------------------------------------------------------------------
func kasusBukuTelepon() {
	telepon := map[string]string{
		"Budi": "08123456789",
		"Siti": "08234567890",
		"Agus": "08345678901",
	}
	fmt.Println("Buku telepon:", telepon)
	nama := "Siti"
	if no, ok := telepon[nama]; ok {
		fmt.Printf("Nomor %s: %s\n", nama, no)
	} else {
		fmt.Printf("%s tidak ada di buku telepon.\n", nama)
	}
}

// -----------------------------------------------------------------------------
// 7. Studi kasus: hitung kemunculan kata (word count)
// -----------------------------------------------------------------------------
func kasusWordCount() {
	kalimat := "satu dua dua tiga tiga tiga"
	jumlah := make(map[string]int)
	for _, kata := range strings.Fields(kalimat) {
		jumlah[kata]++
	}
	fmt.Println("Kalimat:", kalimat)
	fmt.Println("Jumlah per kata:", jumlah)
}

// -----------------------------------------------------------------------------
// 8. Studi kasus: konfigurasi (key → value string)
// -----------------------------------------------------------------------------
func kasusKonfigurasi() {
	config := map[string]string{
		"host": "localhost",
		"port": "8080",
		"env":  "development",
	}
	fmt.Println("Config:", config)
	fmt.Println("Host:", config["host"])
	config["debug"] = "true"
	fmt.Println("Setelah tambah debug:", config)
}

// -----------------------------------------------------------------------------
// 9. Studi kasus: cache lookup (simulasi cache ID → hasil)
// -----------------------------------------------------------------------------
func kasusCacheLookup() {
	cache := map[int]string{
		1: "hasil-A",
		2: "hasil-B",
	}
	id := 2
	if val, ok := cache[id]; ok {
		fmt.Printf("Cache hit id=%d: %s\n", id, val)
	} else {
		fmt.Printf("Cache miss id=%d, fetch dan simpan...\n", id)
		cache[id] = "hasil-B"
	}
	fmt.Println("Cache sekarang:", cache)
}
