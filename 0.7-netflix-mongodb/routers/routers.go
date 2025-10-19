package routers

import (
	"github.com/gorilla/mux"
	"main.go/controllers"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateNewMovie).Methods("POST")
	router.HandleFunc("/api/movie/{movieId}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{movieId}", controllers.DeleteMyOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies", controllers.DeleteMyAllMovies).Methods("DELETE")

	return router
}
