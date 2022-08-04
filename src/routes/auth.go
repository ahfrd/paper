package routes

import (
	"net/http"
	controllers "paper/src/controllers/auths"
	"paper/src/helpers"

	"github.com/gorilla/mux"
)

// HealthCheck is a
func AuthenticationRoutes(r *mux.Router, base string) {
	var uri = helpers.ConRoute
	ControllerRegisterRoute := http.HandlerFunc(controllers.ControllerStructure{}.ControllerRegister)
	r.Handle(uri(base, "/register"), ControllerRegisterRoute).Methods("POST")
	DetailUser := http.HandlerFunc(controllers.ControllerStructure{}.GetDetailUser)
	r.Handle(uri(base, "/get"), DetailUser).Methods("POST")
	ControllerLoginRoute := http.HandlerFunc(controllers.ControllerStructure{}.ControllerLogin)
	r.Handle(uri(base, "/login"), ControllerLoginRoute).Methods("POST")
	ControllerLogoutRoute := http.HandlerFunc(controllers.ControllerStructure{}.ControllerLogout)
	r.Handle(uri(base, "/logout"), ControllerLogoutRoute).Methods("POST")
}
