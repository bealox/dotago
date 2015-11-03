package api

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Name(route.Name).Path(route.Path).Handler(route.Handler)
	}
	return router
}
