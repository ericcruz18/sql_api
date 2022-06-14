package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	//Ping database
	bd, err := getDB()
	if err != nil {
		log.Printf("Error con BD" + err.Error())
		return
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error de conexion a BD. Por favor revisar credenciales. El eror es:" + err.Error())
			return
		}
	}
	//Define rutas
	router := mux.NewRouter()
	SetupRoutesForVideoGames(router)

	port := ":8000"

	server := &http.Server{
		Handler: router,
		Addr:    port,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server iniciando en %s", port)
	log.Fatal(server.ListenAndServe())
}
