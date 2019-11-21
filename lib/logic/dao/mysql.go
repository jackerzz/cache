package dao

import (
	"database/sql"
	"other/gim/conf"
)

var DBCli *sql.DB

func init() {
	var err error
	DBCli, err = sql.Open("mysql", conf.MySQL)
	if err != nil {
		panic(err)
	}
}