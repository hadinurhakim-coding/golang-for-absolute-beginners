package akun

// Username diekspor (huruf besar).
type Profil struct {
	Username string
	password string // tidak diekspor — hanya bisa diakses dari paket akun
}

func BuatProfil(user, pass string) Profil {
	return Profil{Username: user, password: pass}
}

func (p Profil) Ringkasan() string {
	// password bisa dibaca di sini karena masih paket akun
	return p.Username + " (rahasia disembunyikan)"
}
