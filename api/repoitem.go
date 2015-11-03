package api

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// type ItemStruct struct {
// 	ID       int
// 	Name     string    `json:"name"`
// 	DotaID   int       `json:"id" db:"dotaID"`
// 	DotaName string    `json:"dota_name" db:"dotaName"`
// 	Created  time.Time `json:"created" db:"created"`
// }
type ItemStruct struct {
	ID         int       `json:"-"`
	Name       string    `json:"name"`
	DotaID     int       `json:"id"`
	DotaName   string    `json:"dota_name"`
	Cost       int       `json:"cost"`
	SecretShop int       `json:"secret_shop"`
	SideShop   int       `json:"side_shop"`
	Recipe     int       `json:"recipe"`
	Created    time.Time `json:"created"`
}

func GetItem() ([]ItemStruct, error) {
	var db *sqlx.DB
	var err error
	var items []ItemStruct

	if db, err = sqlx.Open("mysql", user+":"+password+"@/Dota?parseTime=true"); err != nil {
		return items, err
	}
	defer db.Close()

	rows, err := db.Queryx("select * from Item")
	if err != nil {
		return items, err
	}

	for rows.Next() {
		var item ItemStruct

		err = rows.StructScan(&item)
		if err != nil {
			return items, err
		}

		items = append(items, item)
	}
	return items, nil
}
