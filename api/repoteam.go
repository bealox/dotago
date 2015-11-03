package api

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type TeamStruct struct {
	ID             int       `json:"-"`
	Name           string    `json:"name"`
	Tag            string    `json:"tag"`
	SteamCreatedBy time.Time `json:"created_by"`
	Logo           uint64    `json:"logo"`
	FileName       string    `json:"-"`
	TeamID         uint64    `json:"team_id"`
	CountryCode    string    `json:"country_code"`
	URL            string    `json:"-"`
	Modified       time.Time `json:"modified"`
	Player0        uint64    `json:"player0"`
	Player1        uint64    `json:"player1"`
	Player2        uint64    `json:"player2"`
	Player3        uint64    `json:"player3"`
	Player4        uint64    `json:"player4"`
	Player5        uint64    `json:"substitute"`
}

func GetTeam() ([]TeamStruct, error) {
	var db *sqlx.DB
	var err error
	var teams []TeamStruct

	if db, err = sqlx.Open("mysql", user+":"+password+"@/Dota?parseTime=true"); err != nil {
		return teams, err
	}
	defer db.Close()

	rows, err := db.Queryx("select * from Team")
	if err != nil {
		return teams, err
	}

	for rows.Next() {
		var team TeamStruct

		err = rows.StructScan(&team)
		if err != nil {
			return teams, err
		}

		teams = append(teams, team)
	}
	return teams, nil
}
