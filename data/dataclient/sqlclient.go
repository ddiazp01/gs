package dataclient

import (
	"database/sql"
	"fmt"
	"gs/data/model"

	_ "github.com/go-sql-driver/mysql" ///El driver se registra en database/sql en su función Init(). Es usado internamente por éste
)

//InsertarPeticion funcion de peticion
func InsertarPeticion(objeto *model.Usuario) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/gs")

	if err != nil {
		panic(err.Error()) //si se abre bien
	}

	defer db.Close() //cerrar la conexion nosotros. Hay que cerrarlo siempre

	insert, err := db.Query("INSERT INTO Usuario(name, apellidos, username, password, email) VALUES (?, ?, ?, ?, ?)", objeto.Name, objeto.Apellidos, objeto.UserName, objeto.Password, objeto.Email)
	//Inserta una nueva peticion en la base de datos,guardar fechas de horarios en utc.
	if err != nil {
		panic(err.Error())
	}
	insert.Close()

}

//LogearUsuario logear usuario
func LogearUsuario(objeto *model.Login) string {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/gs")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Consultamos todos los idiomas de la base de datos
	comando := "SELECT Password FROM Usuario WHERE (UserName = '" + objeto.UserName + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT Password FROM Usuario WHERE (UserName = '" + objeto.UserName + "')")

	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	var resultado string
	for query.Next() {
		err = query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultado
}

//ConsultaID test
/*func ConsultaID(username string) int {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/gs")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	comando := "SELECT ID FROM Usuario WHERE (UserName = '" + username + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT ID FROM Usuario WHERE (UserName = '" + username + "')")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	var resultado int
	for query.Next() {
		err := query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultado
}*/
