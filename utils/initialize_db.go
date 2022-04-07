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
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB, err := gorm.Open(postgres.Open(Db_conn_str), &gorm.Config{})

func Set_db(Db_conn_str string) *gorm.DB {
	DB, err := gorm.Open(postgres.Open(Db_conn_str), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}
	return DB
}
