// =============================================================================
// SWITCH STATEMENT DI GO
// =============================================================================
// switch mengevaluasi satu nilai lalu bandingkan dengan banyak case.
// Bentuk: switch expr { case v1: ... case v2: ... default: ... }
// Di Go: tidak perlu break; setelah case selesai otomatis keluar (no fall-through).
// Pakai fallthrough jika ingin eksekusi lanjut ke case berikutnya.
// =============================================================================

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("========== 1. switch dasar (nilai ekspresi) ==========")
	switchDasar()

	fmt.Println("\n========== 2. switch tanpa ekspresi (seperti if-else if) ==========")
	switchTanpaEkspresi()

	fmt.Println("\n========== 3. switch dengan short statement ==========")
	switchShortStatement()

	fmt.Println("\n========== 4. case gabungan (beberapa nilai) ==========")
	switchCaseGabungan()

	fmt.Println("\n========== 5. fallthrough (lanjut ke case berikutnya) ==========")
	switchFallthrough()

	fmt.Println("\n========== 6. Studi kasus: hari dalam seminggu ==========")
	kasusHari()

	fmt.Println("\n========== 7. Studi kasus: menu aplikasi ==========")
	kasusMenu()

	fmt.Println("\n========== 8. Studi kasus: kategori nilai (grade) ==========")
	kasusGrade()

	fmt.Println("\n========== 9. Studi kasus: tipe pembayaran ==========")
	kasusTipePembayaran()

	fmt.Println("\n========== 10. Studi kasus: musim berdasarkan bulan ==========")
	kasusMusim()
}

// -----------------------------------------------------------------------------
// 1. switch nilai { case v: ... }. Case yang cocok dijalankan; lalu keluar.
// -----------------------------------------------------------------------------
func switchDasar() {
	hari := 3
	switch hari {
	case 1:
		fmt.Println("  Senin")
	case 2:
		fmt.Println("  Selasa")
	case 3:
		fmt.Println("  Rabu")
	case 4:
		fmt.Println("  Kamis")
	default:
		fmt.Println("  Bukan 1-4")
	}
}

// -----------------------------------------------------------------------------
//  2. switch { case kondisi1: ... case kondisi2: ... } — tanpa nilai di switch
//     Setiap case adalah kondisi bool (seperti if-else if).
//
// -----------------------------------------------------------------------------
func switchTanpaEkspresi() {
	nilai := 76
	switch {
	case nilai >= 90:
		fmt.Println("  Grade A")
	case nilai >= 80:
		fmt.Println("  Grade B")
	case nilai >= 70:
		fmt.Println("  Grade C")
	case nilai >= 60:
		fmt.Println("  Grade D")
	default:
		fmt.Println("  Grade E")
	}
}

// -----------------------------------------------------------------------------
// 3. switch x := ...; x { ... } — short statement; x hanya dalam scope switch
// -----------------------------------------------------------------------------
func switchShortStatement() {
	switch jam := time.Now().Hour(); {
	case jam < 12:
		fmt.Println("  Pagi")
	case jam < 18:
		fmt.Println("  Sore")
	default:
		fmt.Println("  Malam")
	}
}

// -----------------------------------------------------------------------------
// 4. Satu case bisa punya beberapa nilai: case 1, 2, 3:
// -----------------------------------------------------------------------------
func switchCaseGabungan() {
	bulan := 2
	switch bulan {
	case 1, 2, 3:
		fmt.Println("  Kuartal 1")
	case 4, 5, 6:
		fmt.Println("  Kuartal 2")
	case 7, 8, 9:
		fmt.Println("  Kuartal 3")
	case 10, 11, 12:
		fmt.Println("  Kuartal 4")
	default:
		fmt.Println("  Bulan tidak valid")
	}
}

// -----------------------------------------------------------------------------
// 5. fallthrough: lanjut eksekusi ke case berikutnya (jarang dipakai)
// -----------------------------------------------------------------------------
func switchFallthrough() {
	x := 1
	switch x {
	case 1:
		fmt.Println("  case 1")
		fallthrough
	case 2:
		fmt.Println("  case 2 (atau fallthrough dari 1)")
	default:
		fmt.Println("  default")
	}
}

// -----------------------------------------------------------------------------
// 6. Studi kasus: angka ke nama hari
// -----------------------------------------------------------------------------
func kasusHari() {
	noHari := 5
	var namaHari string
	switch noHari {
	case 1:
		namaHari = "Senin"
	case 2:
		namaHari = "Selasa"
	case 3:
		namaHari = "Rabu"
	case 4:
		namaHari = "Kamis"
	case 5:
		namaHari = "Jumat"
	case 6:
		namaHari = "Sabtu"
	case 7:
		namaHari = "Minggu"
	default:
		namaHari = "Tidak valid"
	}
	fmt.Printf("  Nomor %d = %s\n", noHari, namaHari)
}

// -----------------------------------------------------------------------------
// 7. Studi kasus: pilihan menu (1=Input, 2=Lihat, 3=Keluar)
// -----------------------------------------------------------------------------
func kasusMenu() {
	pilihan := 2
	switch pilihan {
	case 1:
		fmt.Println("  Menu: Input data")
	case 2:
		fmt.Println("  Menu: Lihat data")
	case 3:
		fmt.Println("  Menu: Keluar")
	default:
		fmt.Println("  Pilihan tidak valid")
	}
}

// -----------------------------------------------------------------------------
// 8. Studi kasus: skor ke grade (A/B/C/D/E)
// -----------------------------------------------------------------------------
func kasusGrade() {
	skor := 88
	var grade string
	switch {
	case skor >= 90:
		grade = "A"
	case skor >= 80:
		grade = "B"
	case skor >= 70:
		grade = "C"
	case skor >= 60:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Printf("  Skor %d → Grade %s\n", skor, grade)
}

// -----------------------------------------------------------------------------
// 9. Studi kasus: tipe pembayaran (transfer, tunai, kartu)
// -----------------------------------------------------------------------------
func kasusTipePembayaran() {
	tipe := "transfer"
	var pesan string
	switch tipe {
	case "transfer":
		pesan = "Masukkan nomor rekening"
	case "tunai":
		pesan = "Siapkan uang cash"
	case "kartu":
		pesan = "Gesek kartu debit/kredit"
	default:
		pesan = "Metode tidak dikenali"
	}
	fmt.Printf("  Tipe %q → %s\n", tipe, pesan)
}

// -----------------------------------------------------------------------------
// 10. Studi kasus: bulan → musim (Des-Jan-Feb=dingin, Mar-Mei=semi, dll)
// -----------------------------------------------------------------------------
func kasusMusim() {
	bulan := 4
	var musim string
	switch bulan {
	case 12, 1, 2:
		musim = "Dingin"
	case 3, 4, 5:
		musim = "Semi"
	case 6, 7, 8:
		musim = "Panas"
	case 9, 10, 11:
		musim = "Gugur"
	default:
		musim = "Tidak valid"
	}
	fmt.Printf("  Bulan %d → Musim %s\n", bulan, musim)
}
