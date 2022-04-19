package config

import "os"

var Db_conn_str = os.Getenv("DATABASE_URL")
