package routes

import (
	"github.com/EDLadder/Hats-For-Parties/controllers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(client *mongo.Client) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/party", controllers.GetParties(client)).Methods("GET")
	router.HandleFunc("/hat", controllers.GetHats(client)).Methods("GET")
	router.HandleFunc("/party/start", controllers.CreateParty(client)).Methods("POST")
	router.HandleFunc("/party/stop/{id}", controllers.StopParty(client)).Methods("PATCH")
	return router
}
