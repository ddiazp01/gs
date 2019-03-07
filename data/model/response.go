package model

import websocket "websocket-master"

// RUsuario struct
type RUsuario struct {
	Name      string
	Apellidos string
	UserName  string
	Password  string
	Email     string
	conn      *websocket.Conn
	send      chan []byte
}

//RLogin struct
type RLogin struct {
	UserName string
	Password string
}
