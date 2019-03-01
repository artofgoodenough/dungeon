package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"
)

type InsertParameters struct {
	branch      string
	label       string
	description string
	tags        []string
	created     int64
	modified    int64
	origin      string
	frozen      bool
}

var Db *sql.DB

func init() {
	var err error
	log.Println("Opening PostgreSQL Database ....")
	Db, err = sql.Open("postgres", "user=ronald dbname=pod_definitions password=reagan sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	log.Println("PostgreSQL Database opened")
}

func (post *InsertParameters) AddRecord() (err error) {
	statement := `
  INSERT INTO Models_Branches_0001 (
  	label,
  	description,
  	tags,
  	created,
  	modified,
  	origin,
  	frozen
  )
  VALUES ($1, $2, $3, $4, $5, $6, $7)
  RETURNING branch;
  `
	log.Printf("query = %s\n", statement)
	stmt, err := Db.Prepare(statement)
	if err != nil {
		log.Panic(err)
		return
	}
	defer stmt.Close()
	log.Print("PostgreSQL Database Insert Statement prepared.")
	err = stmt.QueryRow(
		post.label,
		post.description,
		pq.Array(post.tags),
		post.created,
		post.modified,
		post.origin,
		post.frozen).Scan(&post.branch)
	if err != nil {
		log.Panic(err)
		return
	}
	log.Print("PostgreSQL Database Insert Table completed.")
	return
}

func main() {
	log.Println("Start....")
	insertSql := InsertParameters{branch: "",
		label:       "First",
		description: "My First Entry",
		tags:        []string{"one", "two", "three", "four"},
		created:     time.Now().Unix(),
		modified:    time.Now().Unix(),
		origin:      "{00000000-0000-0000-0000-000000000000}",
		frozen:      false,
	}
	err := insertSql.AddRecord()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Returned: %s\n", insertSql.branch)
	log.Println("Done.")
}
