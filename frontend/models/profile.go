package models

type Profile struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nama     string `json:"nama"`
	NoTelp   string `json:"no_telp"`
}
