package api

import (
	"github.com/bealox/dotahandler/modal"
	"github.com/jmoiron/sqlx"
)

func GetLeague() ([]modal.Modal, error) {
	var db *sqlx.DB
	var err error
	var leagues []modal.Modal

	if db, err = sqlx.Open("mysql", user+":"+password+"@/Dota?parseTime=true"); err != nil {
		return leagues, err
	}
	defer db.Close()

	// find league where deleted is 0000-00-00 00:00:00
	rows, err := db.Queryx("select * from League where deleted = ?", deletedTime)
	if err != nil {
		return leagues, err
	}

	for rows.Next() {
		var league modal.League

		err = rows.StructScan(&league)
		if err != nil {
			return leagues, err
		}

		leagues = append(leagues, league)
	}
	return leagues, nil
}
