package model

import (
	"time"
)

//Usuario struct
type Usuario struct {
	Name      string
	Apellidos string
	UserName  string
	Password  string
	Email     string
}

//Message struct
type Message struct {
	Fecha time.Time
	Texto string
}

//Chat struct
type Chat struct {
	SesionAbierta bool
	SesionCerrada bool
}

//Citas struct
type Citas struct {
	Texto string
	Fecha time.Time
	Hora  time.Time
}

//RespuestasAutomaticas struct
type RespuestasAutomaticas struct {
	Texto string
	Fecha time.Time
}

//Login struct
type Login struct {
	UserName string
	Password string
}
