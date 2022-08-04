package routes

import (
	"net/http"
	controllers "paper/src/controllers/transaction"
	"paper/src/helpers"

	"github.com/gorilla/mux"
)

// HealthCheck is a
func TransactionRoutes(r *mux.Router, base string) {
	var uri = helpers.ConRoute
	ControllerCreateRoute := http.HandlerFunc(controllers.ControllerStructure{}.ControllerInsertTransaction)
	r.Handle(uri(base, "/create"), ControllerCreateRoute).Methods("POST")
	ControllerUpdateRoute := http.HandlerFunc(controllers.ControllerStructure{}.ControllerUpdateTransaction)
	r.Handle(uri(base, "/update"), ControllerUpdateRoute).Methods("POST")
	ControllerDeleteRoute := http.HandlerFunc(controllers.ControllerStructure{}.ControllerDeleteTransaction)
	r.Handle(uri(base, "/delete"), ControllerDeleteRoute).Methods("DELETE")
	GetDataRoute := http.HandlerFunc(controllers.ControllerStructure{}.GetTransaction)
	r.Handle(uri(base, "/get"), GetDataRoute).Methods("POST")
	GetDataDeleteRoute := http.HandlerFunc(controllers.ControllerStructure{}.ShowDeleteData)
	r.Handle(uri(base, "/restore"), GetDataDeleteRoute).Methods("GET")
	GetSummaryRoute := http.HandlerFunc(controllers.ControllerStructure{}.GetTransactionSummary)
	r.Handle(uri(base, "/summary"), GetSummaryRoute).Methods("POST")
}
