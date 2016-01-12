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
	Route{"Get", "Index", "/", index},
	Route{"Get", "HeroIndex", "/heroes", heroIndex},
	Route{"Get", "ItemIndex", "/items", itemIndex},
	Route{"Get", "AbilityIndex", "/abilities", abilityIndex},
	Route{"Get", "TeamIndex", "/teams", teamIndex},
	Route{"Get", "LeagueIndex", "/leagues", leagueIndex},
	Route{"Get", "LiveIndex", "/lives", liveIndex},
}
