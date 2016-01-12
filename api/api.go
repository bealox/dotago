package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	//Mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/bealox/dotahandler/modal"
)

const deletedTime = "0000-00-00 00:00:00"

func processDataToJSON(w http.ResponseWriter, api []modal.Modal) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(api); err != nil {
		log.Println(err)
	}
}

var (
	user     = os.Getenv("MySQLUser")
	password = os.Getenv("MySQLPassword")
)
