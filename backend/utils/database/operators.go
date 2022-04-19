package database

import (
	"context"
	"d/go/config"
	"d/go/errors"
	"d/go/structs"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect_db() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.Connect(context.Background(), config.Db_conn_str)
	if err != nil {
		fmt.Println("failed to connect database")
		fmt.Println("Error is", err)
		err = errors.New_db_connection_error("failed to connect database")
	} else {
		fmt.Println("Db connected sucessfully")
	}
	return dbpool, err
}

func Query_db(comm string) (pgx.Row, error) {
	var row pgx.Row
	Db, err := Connect_db()
	if err != nil {
		fmt.Println("failed to connect database")
		return row, err
	} else {
		defer Db.Close()
		row := Db.QueryRow(context.Background(), comm)
		return row, err
	}

}

func Create_basic_tables() {
	var commnads []string
	commnads = append(commnads, Create_articles_table_command, Create_recently_loaded_articles_table_command, Create_users_table_command)
	for _, comm := range commnads {
		res, err := Query_db(comm)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
	}
}

func Write_article_to_db(article structs.Article_create) string {
	Db, err := Connect_db()
	if err != nil {
		fmt.Println("failed to connect database")
		return "fail"
	} else {
		defer Db.Close()
		insert_string_to_artcls := `insert into articles(title, content, pub_date, image_url, src_link, site_alias)
					  			values(` + `'` + article.Title + `','` + article.Content + `','` + article.Pub_date + `','` + article.Image_url + `','` + article.Src_link + `','` + article.Site_alias + `')`
		insert_string_to_recent := `insert into recently_loaded_articles(pub_date, src_link, site_alias)
								values('` + article.Pub_date + `','` + article.Src_link + `','` + article.Site_alias + `')`
		// delete_string_to_recent := `delete from recently_loaded_articles where site_name = '` + article.Site_alias + `'
		// 													   and pub_date = (select min(pub_date) from recently_loaded_articles)`
		_, err := Db.Query(context.Background(), insert_string_to_artcls)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Inserted article with link to articles " + article.Src_link)
		}
		_, err = Db.Query(context.Background(), insert_string_to_recent)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Inserted article with link to recently loaded" + article.Src_link)
		}
		// rows, err = Db.Query(delete_string_to_recent)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		return "suceeded!"
	}
}

func Select_all_articles() []structs.Article_select {
	Db, err := Connect_db()
	if err != nil {
		fmt.Println("failed to connect database")
	}
	defer Db.Close()
	var articles []structs.Article_select
	// err = pgxscan.Select(context.Background(), Db, &articles, Select_all_articles_command)
	rows, err := Db.Query(context.Background(), Select_all_articles_command)
	if err != nil {
		fmt.Println("1: ")
		fmt.Println(err)
	} else {
		for rows.Next() {
			var a structs.Article_select
			err = rows.Scan(&a.Id, &a.Title, &a.Content, &a.Pub_date, &a.Image_url, &a.Src_link, &a.Site_alias)
			if err != nil {
				fmt.Println(err)
			}
			articles = append(articles, a)
		}
	}
	return articles
}

func Select_article(id string) (structs.Article_select, error) {
	var a structs.Article_select
	Db, err := Connect_db()
	if err != nil {
		fmt.Println(err)
		return a, errors.New_db_connection_error("Failed to connect to db")
	}
	defer Db.Close()

	row := Db.QueryRow(context.Background(), Select_article_command, id)
	err = row.Scan(&a.Id, &a.Title, &a.Content, &a.Pub_date, &a.Image_url, &a.Src_link, &a.Site_alias)
	if err != nil {
		fmt.Println(err)
		if err.Error() == "no rows in result set" {
			return a, errors.Get_api_article_error("No article with such id found!")
		}
	}
	return a, err
}

func Delete_article(id string) (structs.Article_select, error) {
	Db, err := Connect_db()
	var a structs.Article_select
	if err != nil {
		fmt.Println(err)
		return a, errors.New_db_connection_error("Failed to connect to db")
	}
	defer Db.Close()

	row := Db.QueryRow(context.Background(), Delete_article_command, id)
	err = row.Scan(&a.Id, &a.Title, &a.Content, &a.Pub_date, &a.Image_url, &a.Src_link, &a.Site_alias)
	if err != nil {
		fmt.Println(row)
		if err.Error() == "no rows in result set" {
			return a, errors.Get_api_article_error("No article with such id found!")
		}
	}
	return a, err
}

func Create_test_articles() {
	for i := 0; i < 3; i += 1 {
		article := structs.Article_create{
			Title:      "Test article number " + strconv.Itoa(i),
			Content:    "Test text for article number " + strconv.Itoa(i),
			Image_url:  "Test image url for article number " + strconv.Itoa(i),
			Pub_date:   "2022-01-01 17:01:17",
			Src_link:   "test.com",
			Site_alias: "test",
		}
		Write_article_to_db(article)
	}
}

// const Update_article_command = "update articles set title = $1, content = $2, pub_date = $3, image_url = $5, src_link = $6, site_alias = $7"
func Update_article(id string, article structs.Article_create) error {
	Db, err := Connect_db()
	if err != nil {
		fmt.Println(err)
		return errors.New_db_connection_error("Failed to connect to db")
	}
	defer Db.Close()
	command := Update_article_command
	if article.Title != "" {
		command += fmt.Sprintf("title = '%s', ", article.Title)
	}
	if article.Content != "" {
		command += fmt.Sprintf("content = '%s', ", article.Content)
	}
	if article.Pub_date != "" {
		command += fmt.Sprintf("pub_date = '%s', ", article.Pub_date)
	}
	if article.Site_alias != "" {
		command += fmt.Sprintf("site_alias = '%s', ", article.Site_alias)
	}
	if article.Src_link != "" {
		command += fmt.Sprintf("src_link = '%s', ", article.Src_link)
	}
	if article.Image_url != "" {
		command += fmt.Sprintf("image_url = '%s', ", article.Src_link)
	}
	command = command[0 : len(command)-2]
	command += fmt.Sprintf(" where article_id = %s;", id)
	fmt.Println(command)
	Db.QueryRow(context.Background(), command)

	return err
}
