package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	log.Println("Opening PostgreSQL Database ....")
	Db, err = sql.Open("postgres", "user=dummy dbname=dummy password=dummy sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	log.Println("PostgreSQL Database opened")
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "me"}
	log.Println(post)
	post.Create()
	log.Println(post)
	post2, err := GetPost(post.Id)
	if err != nil {
		log.Panic(err)
	}
	log.Println(post2)
}
