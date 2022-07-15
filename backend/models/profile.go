package models

type Login struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Profile struct {
	Nama     string `json:"nama"`
	Password string `json:"password"`
	Email    string `json:"email"`
	NoTelp   string `json:"no_telp"`
}
