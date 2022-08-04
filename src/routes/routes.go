package routes

import (
	"github.com/gorilla/mux"
)

// Route is as
func Route() *mux.Router {

	r := mux.NewRouter()
	AuthenticationRoutes(r, "/auth")
	AccountRoutes(r, "/account")
	TransactionRoutes(r, "/transaction")
	return r
}
