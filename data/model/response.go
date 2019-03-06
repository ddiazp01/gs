package model

import (
	"time"
)

// RUsuario struct
type RUsuario struct {
	ID        int
	Name      string
	Apellidos string
	UserName  string
	Password  string
	Email     string
}

//RLogin struct
type RLogin struct {
	UserName string
	Password string
}

//RCitas struct
type RCitas struct {
	ID    int
	Texto string
	Fecha time.Time
	Hora  time.Time
}
