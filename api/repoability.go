package api

import "github.com/jmoiron/sqlx"

type AbilityStruct struct {
	ID     int    `json:"-"`
	Name   string `json:"name"`
	DotaID int    `json:"id"`
}

func GetAbility() ([]AbilityStruct, error) {
	var db *sqlx.DB
	var err error
	var abilities []AbilityStruct

	if db, err = sqlx.Open("mysql", user+":"+password+"@/Dota?parseTime=true"); err != nil {
		return abilities, err
	}
	defer db.Close()

	rows, err := db.Queryx("select * from Ability")
	if err != nil {
		return abilities, err
	}

	for rows.Next() {
		var ability AbilityStruct

		err = rows.StructScan(&ability)
		if err != nil {
			return abilities, err
		}

		abilities = append(abilities, ability)
	}
	return abilities, nil
}
