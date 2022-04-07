// package utils

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// func Set_db(Db_conn_str string) *sql.DB {
// 	db, err := sql.Open("postgres", Db_conn_str)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Подключение к базе данных было успешно")
// 	}
// 	// defer db.Close()
// 	return db
// }
package utils

import (
	"context"
	"d/go/config"
	"d/go/errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// DB, err := gorm.Open(postgres.Open(Db_conn_str), &gorm.Config{})

func Set_db() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.Connect(context.Background(), config.Db_conn_str)
	if err != nil {
		fmt.Println("failed to connect database")
		err = errors.New_db_connection_error("failed to connect database")
	} else {
		fmt.Println("Db connected sucessfully")
	}
	return dbpool, err
}
