package structs

type Article_create struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Pub_date   string `json:"pub_date"`
	Image_url  string `json:"image_url"`
	Src_link   string `json:"src_link"`
	Site_alias string `json:"site_alias"`
}
