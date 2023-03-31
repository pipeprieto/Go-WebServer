package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct{
	Name 	string
	Email 	string
	Edad 	int
}

func Index(rw http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("template/index.html"))

	user:= User{"Andres Prieto","pipeprieto0902@gmail.com", 22}

	tmpl.Execute(rw,user)

}

func main(){ 
	route:= mux.NewRouter()
	route.HandleFunc("/",Index) //Ruta principal
	//Levantando el servidor
	fmt.Println("Ejecutando el servidor en el puerto 3000")
	http.ListenAndServe("localhost:3000",route)
}

