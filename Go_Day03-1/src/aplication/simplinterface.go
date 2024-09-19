package main

import (
	"log"
	"net/http"
	"src/db"
	"src/handlers"
)

func main() {
	var index handlers.DataBase = "places"
	db.IndexPlaces(index.Name())
	var inter handlers.Store = index

	http.HandleFunc("/", handlers.GetAllRestorans(inter))
	http.HandleFunc("/api/places", handlers.ApiHandler(inter))
	http.HandleFunc("/api/recommend", handlers.CheckJWT("key", handlers.GetNearApi(inter)))
	http.HandleFunc("/api/get_token", handlers.GetJWT("key"))
	log.Fatal(http.ListenAndServe(":8888", nil))

}
