package main

import "fmt"

func main() {
	stokTersedia := 100
	jumlahDiminta := 50
	statusAkunAktif := true
	isBarangRusak := false
	apakahMemberVip := true
	memilikiIzin := true

	isStokCukup := stokTersedia >= jumlahDiminta
	bolehProses := isStokCukup && statusAkunAktif && !isBarangRusak && memilikiIzin

	punyaAksesAman := apakahMemberVip && bolehProses
	siapDikirim := !isBarangRusak

	fmt.Println("=== SISTEM VERIFIKASI PENGIRIMAN ===")
	fmt.Printf("Status Stok : %t\n", isStokCukup)
	fmt.Printf("Status Akun : %t\n", statusAkunAktif)
	fmt.Printf("Status Barang Rusak : %t\n", isBarangRusak)
	fmt.Printf("Status Member Vip : %t\n", apakahMemberVip)
	fmt.Printf("Status Memiliki Izin : %t\n", memilikiIzin)
	fmt.Printf("Status Boleh Proses : %t\n", bolehProses)
	fmt.Printf("Status Punya Akses Aman : %t\n", punyaAksesAman)
	fmt.Printf("Status Siap Dikirim : %t\n", siapDikirim)
	fmt.Println("===================================")

	if punyaAksesAman {
		fmt.Println("Silahkan Proses Pengiriman.")
	} else {
		fmt.Println("Akses tidak aman, perlu verifikasi ulang.")
	}

	if !siapDikirim {
		fmt.Println("Barang rusak, perlu diperbaiki.")
	}

}

//
// CATATAN OPERASI BOOLEAN DI GO
//
// Operator: && (AND), || (OR), ! (NOT). Hasil selalu bool (true/false).
// Short-circuit: evaluasi berhenti begitu hasil pasti — false && ... tidak mengecek sisanya; true || ... juga tidak.
//
// Naming: variabel boolean sering pakai awalan is/has/can (isStokCukup, bolehProses) agar maksud jelas.
//
// Format cetak: %t untuk bool. \n untuk newline di akhir Printf.
//
// Lihat juga: data-type-boolean, operasi-perbandingan, operasi-matematika, if-statement.
//
