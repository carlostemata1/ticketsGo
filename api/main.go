package main

import (
	"database/sql"
	"net/http"
	"fmt"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "xQ47YG6Gtv"
	Contrasenia := "you9gXSOqw"
	Nombre := "xQ47YG6Gtv"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(remotemysql.com)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/medicamentos", Inicio)
	http.HandleFunc("/crearMedicamento", CrearMedicamento)
	http.HandleFunc("/insertarMedicamento", InsertarMedicamento)

	fmt.Println("servidor corriendo")
	http.ListenAndServe(":8000", nil)
}

//---------------------------------------parte del medicamento --------------------------------------
type Medicamento struct {
	Id        int
	Nombre    string
	Precio    float64
	Ubicacion string
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionBD()
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM Medicamento")

	if err != nil {
		panic(err.Error())
	}
	medicamento := Medicamento{}
	arregloMedicameto := []Medicamento{}

	for obtenerRegistros.Next() {
		var id int
		var nombre string
		var precio float64
		var ubicacion string
		err = obtenerRegistros.Scan(&id, &nombre, &precio, &ubicacion)
		if err != nil {
			panic(err.Error())
		}
		medicamento.Id = id
		medicamento.Nombre = nombre
		medicamento.Precio = precio
		medicamento.Ubicacion = ubicacion

		arregloMedicameto = append(arregloMedicameto, medicamento)
	}
	plantillas.ExecuteTemplate(w, "medicamentos", arregloMedicameto)
}

func CrearMedicamento(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crearMedicamento", nil)
}

func InsertarMedicamento(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		precio := r.FormValue("precio")
		ubicacion := r.FormValue("ubicacion")

		conexionEstablecida := conexionBD()
		insertarMedicamento, err := conexionEstablecida.Prepare("INSERT INTO Medicamento (nombre, precio, ubicacion) VALUES (?, ?, ?);")
		if err != nil {
			panic(err.Error())
		}
		insertarMedicamento.Exec(nombre, precio, ubicacion)
		http.Redirect(w, r, "/medicamentos", 301)
	}
}

//---------------------------------------fin del medicamento --------------------------------------
