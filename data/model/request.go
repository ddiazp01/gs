package model

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
