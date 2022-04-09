package config

import "d/go/utils"

// const Db_conn_str string = "user=postgres password=1234 dbname=main_db sslmode=disable"
const Db_conn_str string = "host=localhost user=postgres password=1234 dbname=main_db port=5432 sslmode=disable"

var Sessions = map[string]utils.Session{}
