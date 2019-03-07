package model

import "time"

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

//RPeticion truct
type RPeticion struct {
	ID     int
	Nombre string
	Fecha  time.Time
}
