package model

import "time"

//Usuario struct
type Usuario struct {
	Name      string
	Apellidos string
	UserName  string
	Password  string
	Email     string
}

//Login struct
type Login struct {
	UserName string
	Password string
}

//Peticion struct
type Peticion struct {
	Nombre string
	Fecha  time.Time
}

//Filtro struct
type Filtro struct {
	Fecha time.Time
}
