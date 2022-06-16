package repository

import (
	"time"
)

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	Token     string    `db:"token"`
}

type ProfilUser struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Gender    string    `db:"gender"`
	Birthdate time.Time `db:"birthdate"`
	Address   string    `db:"address"`
	NoHp      int64     `db:"nohp"`
	Instansi  string    `db:"instansi"`
	NoInduk   int64     `db:"noinduk"`
	NamaWali  string    `db:"namawali"`
	GajiOrtu  string    `db:"gajiortu"`
	UserID    int64     `db:"user_id"`
}

type Profil struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Address  string `db:"address"`
	NoHp     int64  `db:"nohp"`
	Instansi string `db:"instansi"`
	Username string `db:"username"`
	UserID   int64  `db:"user_id"`
}

type Berkas struct {
	ID           int64  `db:"id"`
	File         string `db:"file"`
	Essay        string `db:"essay"`
	ProfilUserID int64  `db:"profiluser_id"`
}

type Article struct {
	ID        int64     `db:"id"`
	Judul     string    `db:"judul"`
	Isi       string    `db:"isi"`
	Category  string    `db:"category"`
	CreatedAt time.Time `db:"created_at"`
	ProfilID  int64     `db:"profil_id"`
}
