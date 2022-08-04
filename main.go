package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"paper/src/helpers"
	"paper/src/routes"

	"go.elastic.co/apm/module/apmhttp"
)

func init() {
	env := helpers.Env{}
	env.StartingCheck()
}

func main() {

	urls := routes.Route()

	http.Handle("/", urls)

	portEnv := os.Getenv("MYPORT")
	port := fmt.Sprintf(":%s", portEnv)
	if err := http.ListenAndServe(port, apmhttp.Wrap(urls)); err != nil {
		log.Fatal(err)
	}
}
