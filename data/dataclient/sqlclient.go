package dataclient

import (
	"database/sql"
	"fmt"
	"gs/data/model"
	"time"

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

//InsertarPeticion2 test
func InsertarPeticion2(objeto *model.Peticion) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@1234@tcp(localhost:3306)/gs")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Query("INSERT INTO Peticion(palabra, fecha) VALUES (?, aaaa-mm-ddThh:mi:ss.mmm)", objeto.Nombre, objeto.Fecha.Format(time.RFC3339))
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//ListarRegistros test
func ListarRegistros(objeto *model.Filtro) []model.RPeticion {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@1234@tcp(localhost:3306)/gs?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT * FROM Peticion WHERE (fecha <= '" + objeto.Fecha.Format(time.RFC3339) + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT * FROM Peticion WHERE (fecha >= ?)", objeto.Fecha.Format(time.RFC3339))

	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	resultado := make([]model.RPeticion, 0)
	for query.Next() {
		var fila = model.RPeticion{}

		err = query.Scan(&fila.ID, &fila.Nombre, &fila.Fecha)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, fila)
	}
	return resultado
}
