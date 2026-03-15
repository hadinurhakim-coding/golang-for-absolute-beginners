// =============================================================================
// IF STATEMENT DI GO
// =============================================================================
// if kondisi { blok } — blok dijalankan hanya jika kondisi true.
// Tidak pakai kurung ( ) untuk kondisi (wajib pakai { } untuk body).
// Bisa pakai short statement: if x := ...; kondisi { }
// =============================================================================

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("========== 1. if dasar ==========")
	ifDasar()

	fmt.Println("\n========== 2. if-else ==========")
	ifElse()

	fmt.Println("\n========== 3. if - else if - else ==========")
	ifElseIfElse()

	fmt.Println("\n========== 4. Short statement: if x := ...; kondisi ==========")
	ifShortStatement()

	fmt.Println("\n========== 5. if bersarang (nested) ==========")
	ifBersarang()

	fmt.Println("\n========== 6. Studi kasus: nilai lulus / remidi ==========")
	kasusNilaiLulus()

	fmt.Println("\n========== 7. Studi kasus: kategori usia ==========")
	kasusKategoriUsia()

	fmt.Println("\n========== 8. Studi kasus: tier diskon belanja ==========")
	kasusTierDiskon()

	fmt.Println("\n========== 9. Studi kasus: validasi login & role ==========")
	kasusLoginRole()

	fmt.Println("\n========== 10. Studi kasus: bilangan ganjil/genap ==========")
	kasusGanjilGenap()
}

// -----------------------------------------------------------------------------
// 1. if kondisi { ... }. Kondisi harus bertipe bool.
// -----------------------------------------------------------------------------
func ifDasar() {
	umur := 20
	if umur >= 18 {
		fmt.Println("  Umur >= 18: dewasa.")
	}

	aktif := true
	if aktif {
		fmt.Println("  Status aktif: true.")
	}
}

// -----------------------------------------------------------------------------
// 2. if kondisi { ... } else { ... }
// -----------------------------------------------------------------------------
func ifElse() {
	nilai := 55
	if nilai >= 60 {
		fmt.Println("  Nilai", nilai, ": Lulus.")
	} else {
		fmt.Println("  Nilai", nilai, ": Tidak lulus.")
	}
}

// -----------------------------------------------------------------------------
// 3. if - else if - else (rantai kondisi)
// -----------------------------------------------------------------------------
func ifElseIfElse() {
	skor := 85
	if skor >= 90 {
		fmt.Println("  Skor", skor, ": Grade A.")
	} else if skor >= 80 {
		fmt.Println("  Skor", skor, ": Grade B.")
	} else if skor >= 70 {
		fmt.Println("  Skor", skor, ": Grade C.")
	} else if skor >= 60 {
		fmt.Println("  Skor", skor, ": Grade D.")
	} else {
		fmt.Println("  Skor", skor, ": Grade E.")
	}
}

// -----------------------------------------------------------------------------
// 4. Short statement: variabel hanya dalam scope if/else
// -----------------------------------------------------------------------------
func ifShortStatement() {
	// x dan err hanya ada di dalam blok if/else
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Println("  Parse berhasil, n =", n)
	} else {
		fmt.Println("  Parse gagal:", err)
	}

	// Contoh lain: inisialisasi sekaligus cek
	if x := 10; x > 5 {
		fmt.Println("  x = 10, x > 5 → true")
	}
}

// -----------------------------------------------------------------------------
// 5. if di dalam if (nested)
// -----------------------------------------------------------------------------
func ifBersarang() {
	login := true
	admin := true
	if login {
		if admin {
			fmt.Println("  Login + Admin: akses penuh.")
		} else {
			fmt.Println("  Login saja: akses terbatas.")
		}
	} else {
		fmt.Println("  Belum login.")
	}
}

// -----------------------------------------------------------------------------
// 6. Studi kasus: nilai lulus (>= 60) atau remidi
// -----------------------------------------------------------------------------
func kasusNilaiLulus() {
	nilai := 75
	batasLulus := 60
	if nilai >= batasLulus {
		fmt.Printf("  Nilai %d >= %d → Lulus.\n", nilai, batasLulus)
	} else {
		fmt.Printf("  Nilai %d < %d → Remidi.\n", nilai, batasLulus)
	}
}

// -----------------------------------------------------------------------------
// 7. Studi kasus: kategori usia (anak, remaja, dewasa, lansia)
// -----------------------------------------------------------------------------
func kasusKategoriUsia() {
	usia := 25
	if usia < 13 {
		fmt.Printf("  Usia %d: Anak.\n", usia)
	} else if usia < 20 {
		fmt.Printf("  Usia %d: Remaja.\n", usia)
	} else if usia < 60 {
		fmt.Printf("  Usia %d: Dewasa.\n", usia)
	} else {
		fmt.Printf("  Usia %d: Lansia.\n", usia)
	}
}

// -----------------------------------------------------------------------------
// 8. Studi kasus: tier diskon berdasarkan total belanja
// -----------------------------------------------------------------------------
func kasusTierDiskon() {
	totalBelanja := 250000
	var diskonPersen int
	if totalBelanja >= 500000 {
		diskonPersen = 20
	} else if totalBelanja >= 300000 {
		diskonPersen = 15
	} else if totalBelanja >= 100000 {
		diskonPersen = 10
	} else {
		diskonPersen = 0
	}
	diskon := totalBelanja * diskonPersen / 100
	bayar := totalBelanja - diskon
	fmt.Printf("  Total: Rp %d → Diskon %d%% = Rp %d → Bayar Rp %d\n", totalBelanja, diskonPersen, diskon, bayar)
}

// -----------------------------------------------------------------------------
// 9. Studi kasus: validasi login dan role (admin/user)
// -----------------------------------------------------------------------------
func kasusLoginRole() {
	username := "admin"
	passwordBenar := true
	isAdmin := true

	if !passwordBenar {
		fmt.Println("  Login gagal: password salah.")
		return
	}
	fmt.Printf("  User %s login berhasil.\n", username)
	if isAdmin {
		fmt.Println("  Role: Admin → akses dashboard admin.")
	} else {
		fmt.Println("  Role: User → akses dashboard user.")
	}
}

// -----------------------------------------------------------------------------
// 10. Studi kasus: cek bilangan ganjil/genap
// -----------------------------------------------------------------------------
func kasusGanjilGenap() {
	n := 7
	if n%2 == 0 {
		fmt.Printf("  %d adalah genap.\n", n)
	} else {
		fmt.Printf("  %d adalah ganjil.\n", n)
	}
}
