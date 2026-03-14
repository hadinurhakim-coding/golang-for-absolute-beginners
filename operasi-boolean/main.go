// =============================================================================
// OPERASI BOOLEAN DI GO
// =============================================================================
// Operator: && (AND), || (OR), ! (NOT)
// Short-circuit: && berhenti di false pertama; || berhenti di true pertama.
// Sering dipakai bersama operasi perbandingan di kondisi if/for.
// =============================================================================

package main

import "fmt"

func main() {
	fmt.Println("========== 1. Dasar: &&, ||, ! ==========")
	dasar()

	fmt.Println("\n========== 2. Tabel kebenaran ==========")
	tabelKebenaran()

	fmt.Println("\n========== 3. Short-circuit ==========")
	shortCircuit()

	fmt.Println("\n========== 4. Gabungan dengan perbandingan ==========")
	gabunganPerbandingan()

	fmt.Println("\n========== 5. Contoh kasus: login & role ==========")
	kasusLoginRole()

	fmt.Println("\n========== 6. Contoh kasus: validasi form ==========")
	kasusValidasiForm()

	fmt.Println("\n========== 7. Contoh kasus: discount & membership ==========")
	kasusDiscount()
}

func dasar() {
	fmt.Println("true && true  =", true && true)
	fmt.Println("true && false =", true && false)
	fmt.Println("true || false =", true || false)
	fmt.Println("false || false =", false || false)
	fmt.Println("!true =", !true)
	fmt.Println("!false =", !false)
}

func tabelKebenaran() {
	fmt.Println("AND (&&): true hanya jika keduanya true")
	fmt.Println("  true  && true  → true")
	fmt.Println("  true  && false → false")
	fmt.Println("  false && true  → false")
	fmt.Println("  false && false → false")
	fmt.Println("OR (||): true jika salah satu true")
	fmt.Println("  true  || true  → true")
	fmt.Println("  true  || false → true")
	fmt.Println("  false || true  → true")
	fmt.Println("  false || false → false")
}

func shortCircuit() {
	// Simulasi: fungsi yang return bool dan punya "efek samping" (cetak)
	call := func(name string, result bool) bool {
		fmt.Printf("  [dipanggil %s] ", name)
		return result
	}
	fmt.Println("false && call(\"B\", true):")
	_ = false && call("B", true)
	fmt.Println("(B tidak dipanggil — short-circuit)")

	fmt.Println("true || call(\"Y\", false):")
	_ = true || call("Y", false)
	fmt.Println("(Y tidak dipanggil — short-circuit)")
}

func gabunganPerbandingan() {
	usia := 20
	punyaKTP := true
	bolehVoting := usia >= 17 && punyaKTP
	fmt.Printf("Usia %d, punya KTP %t → boleh voting? %t\n", usia, punyaKTP, bolehVoting)

	nilai := 45
	remidi := nilai < 50 || nilai > 100
	fmt.Printf("Nilai %d → perlu remidi? %t\n", nilai, remidi)
}

func kasusLoginRole() {
	loggedIn := true
	isAdmin := false
	bisaAksesAdmin := loggedIn && isAdmin
	fmt.Printf("Login: %t, Admin: %t → akses halaman admin? %t\n", loggedIn, isAdmin, bisaAksesAdmin)

	loggedIn2 := true
	isAdmin2 := true
	fmt.Printf("Login: %t, Admin: %t → akses halaman admin? %t\n", loggedIn2, isAdmin2, loggedIn2 && isAdmin2)
}

func kasusValidasiForm() {
	email := "user@mail.com"
	password := "secret123"
	panjangMinPassword := 8
	emailValid := len(email) > 0 && len(email) >= 5
	passwordValid := len(password) >= panjangMinPassword
	formValid := emailValid && passwordValid
	fmt.Printf("Email valid? %t, Password valid? %t → Form valid? %t\n", emailValid, passwordValid, formValid)
}

func kasusDiscount() {
	belanjaMinimal := 100000
	totalBelanja := 150000
	anggotaPremium := true
	// Dapat diskon jika: belanja >= minimal ATAU anggota premium
	dapatDiskon := totalBelanja >= belanjaMinimal || anggotaPremium
	fmt.Printf("Total belanja %d, minimal %d, premium %t → dapat diskon? %t\n",
		totalBelanja, belanjaMinimal, anggotaPremium, dapatDiskon)
}
