package api

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// type HeroStruct struct {
// 	ID       int
// 	Name     string    `json:"name"`
// 	DotaID   int       `json:"id" db:"dotaID"`
// 	DotaName string    `json:"dota_name" db:"dotaName"`
// 	Created  time.Time `json:"created" db:"created"`
// }
type HeroStruct struct {
	ID       int       `json:"-"`
	Name     string    `json:"name"`
	DotaID   int       `json:"id"`
	DotaName string    `json:"dota_name"`
	Created  time.Time `json:"created"`
}

func GetHero() ([]HeroStruct, error) {
	var db *sqlx.DB
	var err error
	var heroes []HeroStruct

	if db, err = sqlx.Open("mysql", user+":"+password+"@/Dota?parseTime=true"); err != nil {
		return heroes, err
	}
	defer db.Close()

	rows, err := db.Queryx("select * from Hero")
	if err != nil {
		return heroes, err
	}

	for rows.Next() {
		var hero HeroStruct

		err = rows.StructScan(&hero)
		if err != nil {
			return heroes, err
		}

		heroes = append(heroes, hero)
	}
	return heroes, nil
}
