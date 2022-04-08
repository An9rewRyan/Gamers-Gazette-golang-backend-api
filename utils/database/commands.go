package database

const Create_articles_table_command = `create table articles (
	article_id serial primary key,
	title varchar(300) not null,
	content text not null, 
	pub_date timestamp not null,
	image_url varchar(300) not null,
	src_link varchar(300) not null,
	site_alias varchar(10) not null
);`

const Create_recently_loaded_articles_table_command = `create table recently_loaded_articles (
	pub_date timestamp not null,
	src_link varchar(300) not null,
	site_name varchar(10) not null
);`

const Select_all_articles_command = `select * from articles;`
const Select_all_new_articles_command = `select * from recently_loaded_articles;`

const Select_article_command = "select * from articles where article_id = $1"
