package routes

import (
	"net/http"
	controllers "paper/src/controllers/account"
	"paper/src/helpers"

	"github.com/gorilla/mux"
)

// HealthCheck is a
func AccountRoutes(r *mux.Router, base string) {
	var uri = helpers.ConRoute
	ControllerCreateRoute := http.HandlerFunc(controllers.ControllerStructure{}.ControllerCreateAccount)
	r.Handle(uri(base, "/create"), ControllerCreateRoute).Methods("POST")
	ControllerUpdateRoute := http.HandlerFunc(controllers.ControllerStructure{}.ControllerUpdateAccount)
	r.Handle(uri(base, "/update"), ControllerUpdateRoute).Methods("POST")
	ControllerDeleteRoute := http.HandlerFunc(controllers.ControllerStructure{}.ControllerDeleteAccount)
	r.Handle(uri(base, "/delete"), ControllerDeleteRoute).Methods("DELETE")
	GetDataRoute := http.HandlerFunc(controllers.ControllerStructure{}.GetAccount)
	r.Handle(uri(base, "/get"), GetDataRoute).Methods("POST")
	GetDataRestore := http.HandlerFunc(controllers.ControllerStructure{}.RestoreDataAccount)
	r.Handle(uri(base, "/restore"), GetDataRestore).Methods("GET")

}
