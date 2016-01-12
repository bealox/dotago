package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/bealox/dotago/api"
)

func main() {
	r := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
