package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Welcome to Dota API. This site has the best API in the world")
}

func HeroIndex(w http.ResponseWriter, r *http.Request) {
	heroes, err := GetHero()
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(heroes); err != nil {
		log.Println(err)
	}
}

func ItemIndex(w http.ResponseWriter, r *http.Request) {
	items, err := GetItem()
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(items); err != nil {
		log.Println(err)
	}
}

func AbilityIndex(w http.ResponseWriter, r *http.Request) {
	abilities, err := GetAbility()
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(abilities); err != nil {
		log.Println(err)
	}
}

func TeamIndex(w http.ResponseWriter, r *http.Request) {
	teams, err := GetTeam()
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(teams); err != nil {
		log.Println(err)
	}
}
