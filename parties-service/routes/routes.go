package routes

import (
	"github.com/EDLadder/Hats-For-Parties/controllers"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/party", controllers.GetParties).Methods("GET")
	router.HandleFunc("/hat", controllers.GetHats).Methods("GET")
	router.HandleFunc("/party/start", controllers.CreateParty).Methods("POST")
	router.HandleFunc("/party/stop/{id}", controllers.StopParty).Methods("PATCH")
	return router
}
