package config

import "os"

// const Db_conn_str string = "user=postgres password=1234 dbname=main_db sslmode=disable"
// const Db_conn_str string = "host=db user=postgres password=1234 dbname=main_db port=5432 sslmode=disable"
var Db_conn_str = os.Getenv("DATABASE_URL")
