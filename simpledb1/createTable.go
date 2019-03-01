package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type CreateStatement struct {
	sqlString string
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

func (query *CreateStatement) Create() (err error) {
	log.Printf("query = %s\n", query)
	if _, err = Db.Exec(query.sqlString); err != nil {
		log.Panic(err)
		return
	}
	log.Print("PostgreSQL Database Table created.")
	return
}

func main() {
	log.Println("Start....")
	tblSql := CreateStatement{sqlString: `
    DROP TABLE IF EXISTS Models_Branches_0001;
    CREATE TABLE Models_Branches_0001 (
    	branch uuid DEFAULT uuid_generate_v4(),
    	label TEXT,
    	description TEXT,
    	tags TEXT[],
    	created INTEGER,
    	modified INTEGER,
    	origin uuid,
    	frozen BOOLEAN,
    	PRIMARY KEY (branch)
    );
  `}
	tblSql.Create()
	log.Println("Done.")
}
