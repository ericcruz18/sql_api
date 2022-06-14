package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func respondWithSuccess(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func repondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-allow-Origin", AllowedCORSDomain)
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length, Accept-Encoding, X-CSRF-Token")
			next.ServeHTTP(w, r)
		})
}

func SetupRoutesForVideoGames(router *mux.Router) {

	enableCORS(router)

	router.HandleFunc("/videogames", func(w http.ResponseWriter, r *http.Request) {
		videoGames, err := getVideoGames()
		if err != nil {
			repondWithError(err, w)
		} else {
			respondWithSuccess(videoGames, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/videogame/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			repondWithError(err, w)
			return
		}
		videoGames, err := getVideoGameById(id)
		if err != nil {
			repondWithError(err, w)
		} else {
			respondWithSuccess(videoGames, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/videogames", func(w http.ResponseWriter, r *http.Request) {
		var videoGame VideoGame
		err := json.NewDecoder(r.Body).Decode(&videoGame)
		if err != nil {
			repondWithError(err, w)
		} else {
			err := createVideoGame(videoGame)
			if err != nil {
				repondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/videogame/{id}", func(w http.ResponseWriter, r *http.Request) {
		var videoGame VideoGame
		err := json.NewDecoder(r.Body).Decode(&videoGame)
		if err != nil {
			repondWithError(err, w)
		} else {
			err := updateVideoGame(videoGame)
			if err != nil {
				repondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)

	router.HandleFunc("/videogames/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			repondWithError(err, w)
			return
		}
		err = deleteVideoGame(id)
		if err != nil {
			repondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)

}
