package api

import (
	"github.com/bealox/dotahandler/modal"
	"github.com/jmoiron/sqlx"
)

func GetAbility() ([]modal.Modal, error) {
	var db *sqlx.DB
	var err error
	var abilities []modal.Modal

	if db, err = sqlx.Open("mysql", user+":"+password+"@/Dota?parseTime=true"); err != nil {
		return abilities, err
	}
	defer db.Close()

	rows, err := db.Queryx("select * from Ability")
	if err != nil {
		return abilities, err
	}

	for rows.Next() {
		var ability modal.Ability

		err = rows.StructScan(&ability)
		if err != nil {
			return abilities, err
		}

		abilities = append(abilities, ability)
	}
	return abilities, nil
}
