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

//InsertarCita funcion de peticion
func InsertarCita(objeto *model.Citas) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/gs")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close() //cerrar la conexion.

	insert, err := db.Query("INSERT INTO Citas(texto, fecha, hora) VALUES (?, ?, ?)", objeto.Texto, objeto.Fecha, objeto.Hora)
	//Inserta una nueva cita en la base de datos,guardar fechas de citas.
	if err != nil {
		panic(err.Error())
	}
	insert.Close()

}

//ConsultarCitas test
func ConsultarCitas() []model.RCitas {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/gs?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT * FROM Citas"
	fmt.Println(comando)
	query, err := db.Query("SELECT * FROM Citas ORDER BY id DESC")

	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	resultado := make([]model.RCitas, 0)
	for query.Next() {
		var fila = model.RCitas{}

		err = query.Scan(&fila.ID, &fila.Texto, &fila.Fecha, &fila.Hora)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, fila)
	}
	return resultado
}

//ActualizarCitas funcion de peticion
func ActualizarCitas(objeto *model.Citas) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/gs")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close() //cerrar la conexion.

	update, err := db.Query("UPDATE Citas SET (texto, fecha, hora) VALUES (?, ?, ?) WHERE ID =?)", objeto.Texto, objeto.Fecha, objeto.Hora)
	//Actulaiza una nueva cita en la base de datos,guardar fechas de horarios en utc.
	if err != nil {
		panic(err.Error())
	}
	update.Close()

}

//EliminarCitas funcion de peticion
func EliminarCitas(objeto *model.Citas) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/gs")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close() //cerrar la conexion.

	delete, err := db.Query("DELETE FROM Citas WHERE ID =?", objeto.Texto, objeto.Fecha, objeto.Fecha)
	//Actulaiza una nueva cita en la base de datos,guardar fechas de horarios en utc.
	if err != nil {
		panic(err.Error())
	}
	delete.Close()

}
