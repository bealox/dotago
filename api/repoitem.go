package api

import (
	"github.com/bealox/dotahandler/modal"
	"github.com/jmoiron/sqlx"
)

// type ItemStruct struct {
// 	ID       int
// 	Name     string    `json:"name"`
// 	DotaID   int       `json:"id" db:"dotaID"`
// 	DotaName string    `json:"dota_name" db:"dotaName"`
// 	Created  time.Time `json:"created" db:"created"`
// }

func GetItem() ([]modal.Modal, error) {
	var db *sqlx.DB
	var err error
	var items []modal.Modal

	if db, err = sqlx.Open("mysql", user+":"+password+"@/Dota?parseTime=true"); err != nil {
		return items, err
	}
	defer db.Close()

	rows, err := db.Queryx("select * from Item")
	if err != nil {
		return items, err
	}

	for rows.Next() {
		var item modal.Item

		err = rows.StructScan(&item)
		if err != nil {
			return items, err
		}

		items = append(items, item)
	}
	return items, nil
}
