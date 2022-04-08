package structs

import "time"

type Article struct {
	Id         int       `json:"article_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Pub_date   time.Time `json:"pub_date"`
	Image_url  string    `json:"image_url"`
	Src_link   string    `json:"src_link"`
	Site_alias string    `json:"site_alias"`
}
