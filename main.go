package main

import (
	"net/http"
	"os"

	"gomerch/helpers"
	"gomerch/routes"
)

func main() {

	if helpers.Contains(os.Args, "--populate") {
		helpers.Populate()
	}

	routes.Config()
	http.ListenAndServe(":8000", nil)
}
