package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := ""
	Nombre := "basedatosprueba"
	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/eliminar", Eliminar)
	http.HandleFunc("/listarTickets", ListarTickets)
	http.HandleFunc("/borrados", Borrados)
	http.HandleFunc("/recuperar", Recuperar)
	http.HandleFunc("/crearTicket", CrearTicket)
	http.HandleFunc("/insertarTicket", InsertarTicket)
	http.HandleFunc("/editar", Editar)
	http.HandleFunc("/actualizar", Actualizar)

	fmt.Println("servidor corriendo")
	http.ListenAndServe(":8000", nil)
}

//---------------------------------------parte del funciones Tickets --------------------------------------
type Ticket struct {
	Id                 int
	Usuario            string
	FechaCreacion      string
	FechaActualizacion string
	Estatus            bool
}

func ListarTickets(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionBD()
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM tickets")

	if err != nil {
		panic(err.Error())
	}
	ticket := Ticket{}
	arregloTicket := []Ticket{}

	for obtenerRegistros.Next() {
		var id int
		var usuario string
		var fechaCreacion string
		var fechaActualizacion string
		var estatus bool
		err = obtenerRegistros.Scan(&id, &usuario, &fechaCreacion, &fechaActualizacion, &estatus)
		if err != nil {
			panic(err.Error())
		}
		ticket.Id = id
		ticket.Usuario = usuario
		ticket.FechaCreacion = fechaCreacion
		ticket.FechaActualizacion = fechaActualizacion
		ticket.Estatus = estatus

		arregloTicket = append(arregloTicket, ticket)
	}
	plantillas.ExecuteTemplate(w, "listarTickets", arregloTicket)
}

func CrearTicket(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crearTicket", nil)
}

func InsertarTicket(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		usuario := r.FormValue("usuario")
		estatus := r.FormValue("estatus")
		estado := 0
		if estatus == "on" {
			estado = 1
		}
		conexionEstablecida := conexionBD()
		insertarTicket, err := conexionEstablecida.Prepare("INSERT INTO tickets (usuario, estatus) VALUES (?,?);")
		if err != nil {
			panic(err.Error())
		}
		insertarTicket.Exec(usuario, estado)
		http.Redirect(w, r, "/listarTickets", 301)
	}
}

func Eliminar(w http.ResponseWriter, r *http.Request) {
	idTicket := r.URL.Query().Get("id")
	conexionEstablecida := conexionBD()

	obtenerRegistro, err := conexionEstablecida.Query("SELECT * FROM tickets WHERE id=?", idTicket)
	if err != nil {
		panic(err.Error())
	}
	ticket := Ticket{}
	for obtenerRegistro.Next() {
		var id int
		var usuario string
		var fechaCreacion string
		var fechaActualizacion string
		var estatus bool
		err = obtenerRegistro.Scan(&id, &usuario, &fechaCreacion, &fechaActualizacion, &estatus)
		if err != nil {
			panic(err.Error())
		}
		ticket.Id = id
		ticket.Usuario = usuario
		ticket.FechaCreacion = fechaCreacion
		ticket.FechaActualizacion = fechaActualizacion
		ticket.Estatus = estatus
		insertarRegistro, err := conexionEstablecida.Prepare("INSERT INTO ticketborrados (usuario, fechaCreacion, fechaActualizacion,estatus) VALUES (?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insertarRegistro.Exec(usuario, fechaCreacion, fechaActualizacion, estatus)
	}

	borrarTicket, err := conexionEstablecida.Prepare("DELETE FROM tickets WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	borrarTicket.Exec(idTicket)
	http.Redirect(w, r, "/listarTickets", 301)

}

func Editar(w http.ResponseWriter, r *http.Request) {
	idTicket := r.URL.Query().Get("id")
	conexionEstablecida := conexionBD()
	obtenerRegistro, err := conexionEstablecida.Query("SELECT * FROM tickets WHERE id=?", idTicket)
	ticket := Ticket{}
	for obtenerRegistro.Next() {
		var id int
		var usuario string
		var fechaCreacion string
		var fechaActualizacion string
		var estatus bool
		err = obtenerRegistro.Scan(&id, &usuario, &fechaCreacion, &fechaActualizacion, &estatus)
		if err != nil {
			panic(err.Error())
		}
		ticket.Id = id
		ticket.Usuario = usuario
		ticket.FechaCreacion = fechaCreacion
		ticket.FechaActualizacion = fechaActualizacion
		ticket.Estatus = estatus
	}

	plantillas.ExecuteTemplate(w, "editar", ticket)

}

func Actualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		usuario := r.FormValue("usuario")
		estatus := r.FormValue("estatus")
		estado := 0
		if estatus == "on" {
			estado = 1
		}
		conexionEstablecida := conexionBD()
		modificarTicket, err := conexionEstablecida.Prepare("UPDATE tickets SET usuario=?,estatus=?,fechaActualizacion=CURRENT_TIMESTAMP WHERE id=?")
		if err != nil {
			panic(err.Error())
		}

		modificarTicket.Exec(usuario, estado, id)
		http.Redirect(w, r, "/listarTickets", 301)
	}
}

func Borrados(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionBD()
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM ticketborrados")

	if err != nil {
		panic(err.Error())
	}
	ticket := Ticket{}
	arregloTicket := []Ticket{}

	for obtenerRegistros.Next() {
		var id int
		var usuario string
		var fechaCreacion string
		var fechaActualizacion string
		var estatus bool
		err = obtenerRegistros.Scan(&id, &usuario, &fechaCreacion, &fechaActualizacion, &estatus)
		if err != nil {
			panic(err.Error())
		}
		ticket.Id = id
		ticket.Usuario = usuario
		ticket.FechaCreacion = fechaCreacion
		ticket.FechaActualizacion = fechaActualizacion
		ticket.Estatus = estatus

		arregloTicket = append(arregloTicket, ticket)
	}
	plantillas.ExecuteTemplate(w, "borrados", arregloTicket)
}

func Recuperar(w http.ResponseWriter, r *http.Request) {
	idTicket := r.URL.Query().Get("id")
	conexionEstablecida := conexionBD()

	obtenerRegistro, err := conexionEstablecida.Query("SELECT * FROM ticketborrados WHERE id=?", idTicket)
	ticket := Ticket{}
	for obtenerRegistro.Next() {
		var id int
		var usuario string
		var fechaCreacion string
		var fechaActualizacion string
		var estatus bool
		err = obtenerRegistro.Scan(&id, &usuario, &fechaCreacion, &fechaActualizacion, &estatus)
		if err != nil {
			panic(err.Error())
		}
		ticket.Id = id
		ticket.Usuario = usuario
		ticket.FechaCreacion = fechaCreacion
		ticket.FechaActualizacion = fechaActualizacion
		ticket.Estatus = estatus
		insertarRegistro, err := conexionEstablecida.Prepare("INSERT INTO tickets (usuario, fechaCreacion, fechaActualizacion,estatus) VALUES (?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insertarRegistro.Exec(usuario, fechaCreacion, fechaActualizacion, estatus)
	}

	borrarTicket, err := conexionEstablecida.Prepare("DELETE FROM ticketborrados WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	borrarTicket.Exec(idTicket)
	http.Redirect(w, r, "/borrados", 301)
}
