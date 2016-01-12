package api

import (
	"github.com/bealox/dotahandler/modal"
	"github.com/jmoiron/sqlx"
)

func GetTeam() ([]modal.Modal, error) {
	var db *sqlx.DB
	var err error
	var teams []modal.Modal

	if db, err = sqlx.Open("mysql", user+":"+password+"@/Dota?parseTime=true"); err != nil {
		return teams, err
	}
	defer db.Close()

	rows, err := db.Queryx("select * from Team")
	if err != nil {
		return teams, err
	}

	for rows.Next() {
		var team modal.Team

		err = rows.StructScan(&team)
		if err != nil {
			return teams, err
		}

		teams = append(teams, team)
	}
	return teams, nil
}
