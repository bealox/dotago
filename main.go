package main

import (
	"net/http"

	log "github.com/bealox/dotaAPI/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/bealox/dotaAPI/api"
)

func main() {
	r := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
