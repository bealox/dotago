package api

import (
	log "github.com/Sirupsen/logrus"
	modal "github.com/bealox/dotahandler/modal"
	"github.com/jmoiron/sqlx"
)

func GetLiveMatch() ([]modal.Modal, error) {
	var db *sqlx.DB
	var err error
	var matches []modal.Modal

	if db, err = sqlx.Open("mysql", user+":"+password+"@/Dota?parseTime=true"); err != nil {
		return matches, err
	}
	defer db.Close()

	rows, err := db.Queryx("select * from Live where deleted=?", deletedTime)
	if err != nil {
		return matches, err
	}

	for rows.Next() {
		var match modal.Live

		err = rows.StructScan(&match)
		if err != nil {
			return matches, err
		}

		log.WithFields(log.Fields{
			"radiant id ": match.RadiantID,
			"dire id ":    match.DireID,
		}).Info("Searching for Details in Score table")

		err = db.QueryRowx("SELECT * FROM Score  where id=?  LIMIT 1", match.RadiantID).StructScan(&match.Radiant)
		if err != nil {
			return matches, err
		}

		_, err = assignDataToPicksAndBans(db, &match)
		if err != nil {
			return matches, err
		}

		matches = append(matches, match)
	}
	return matches, nil
}

func assignDataToPicksAndBans(db *sqlx.DB, match *modal.Live) (matches []modal.Modal, err error) {
	var rpicks []modal.Pick
	err = db.Select(&rpicks, "select t2.dotaid, t2.name from Pick t1, Hero t2 where t1.heroid = t2.dotaid and t1.scoreid=?", match.RadiantID)
	if err != nil {
		return matches, err
	}
	match.Radiant.Picks = append(match.Radiant.Picks, rpicks...)

	var rbans []modal.Ban
	err = db.Select(&rbans, "select t2.dotaid, t2.name from Ban t1, Hero t2 where t1.heroid = t2.dotaid and t1.scoreid=?", match.RadiantID)
	if err != nil {
		return matches, err
	}
	match.Radiant.Bans = append(match.Radiant.Bans, rbans...)

	err = db.QueryRowx("SELECT * FROM Score  where id=? LIMIT 1", match.DireID).StructScan(&match.Dire)
	if err != nil {
		return matches, err
	}

	var dpicks []modal.Pick
	err = db.Select(&dpicks, "select t2.dotaid, t2.name from Pick t1, Hero t2 where t1.heroid = t2.dotaid and t1.scoreid=?", match.DireID)
	if err != nil {
		return matches, err
	}
	match.Dire.Picks = append(match.Dire.Picks, dpicks...)

	var dbans []modal.Ban
	err = db.Select(&dbans, "select t2.dotaid, t2.name from Ban t1, Hero t2 where t1.heroid = t2.dotaid and t1.scoreid=?", match.DireID)
	if err != nil {
		return matches, err
	}
	match.Dire.Bans = append(match.Dire.Bans, dbans...)

	return matches, nil
}
