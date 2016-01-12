package api

import (
	"github.com/bealox/dotahandler/modal"
	"github.com/jmoiron/sqlx"
)

func GetHero() ([]modal.Modal, error) {
	var db *sqlx.DB
	var err error
	var heroes []modal.Modal

	if db, err = sqlx.Open("mysql", user+":"+password+"@/Dota?parseTime=true"); err != nil {
		return heroes, err
	}
	defer db.Close()

	rows, err := db.Queryx("select * from Hero")
	if err != nil {
		return heroes, err
	}

	for rows.Next() {
		var hero modal.Hero

		err = rows.StructScan(&hero)
		if err != nil {
			return heroes, err
		}

		heroes = append(heroes, hero)
	}
	return heroes, nil
}
