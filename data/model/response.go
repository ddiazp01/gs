package model

<<<<<<< HEAD
import (
	"time"
)
=======
import "time"
>>>>>>> 8cda52efc8a5d19aabafe99f889d269cfad83798

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

<<<<<<< HEAD
//RCitas struct
type RCitas struct {
	ID    int
	Texto string
	Fecha time.Time
	Hora  time.Time
=======
//RPeticion truct
type RPeticion struct {
	ID     int
	Nombre string
	Fecha  time.Time
>>>>>>> 8cda52efc8a5d19aabafe99f889d269cfad83798
}
