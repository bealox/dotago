package api

import "net/http"

type Route struct {
	Method  string
	Name    string
	Path    string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	//Get
	Route{"Get", "Index", "/", Index},
	Route{"Get", "HeroIndex", "/heroes", HeroIndex},
	Route{"Get", "ItemIndex", "/items", ItemIndex},
	Route{"Get", "AbilityIndex", "/abilities", AbilityIndex},
	Route{"Get", "TeamIndex", "/teams", TeamIndex},
}
