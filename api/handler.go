package api

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Welcome to Dota 2 API. This site has the best tournament API in the world!")
}

func heroIndex(w http.ResponseWriter, r *http.Request) {
	heroes, err := GetHero()
	if err != nil {
		log.Println(err)
	}
	processDataToJSON(w, heroes)
}

func itemIndex(w http.ResponseWriter, r *http.Request) {
	items, err := GetItem()
	if err != nil {
		log.Println(err)
	}
	processDataToJSON(w, items)
}

func abilityIndex(w http.ResponseWriter, r *http.Request) {
	abilities, err := GetAbility()
	if err != nil {
		log.Println(err)
	}
	processDataToJSON(w, abilities)
}

func teamIndex(w http.ResponseWriter, r *http.Request) {
	teams, err := GetTeam()
	if err != nil {
		log.Println(err)
	}
	processDataToJSON(w, teams)
}

func leagueIndex(w http.ResponseWriter, r *http.Request) {
	leagues, err := GetLeague()
	if err != nil {
		log.Println(err)
	}

	processDataToJSON(w, leagues)
}

func liveIndex(w http.ResponseWriter, r *http.Request) {

	live, err := GetLiveMatch()
	if err != nil {
		log.Println(err)
	}
	processDataToJSON(w, live)
}
