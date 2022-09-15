package mdb

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

type ManagerDb struct {
	Pg *sql.DB
}

var Pg *sql.DB

func (r *ManagerDb) Init() {

	pgHost := os.Getenv("POSTGRES")

	if pgHost == "" {
		log.Fatal("postgres enviroment varible could not find")
	}

	Pg, err := sql.Open("postgres", "postgres://postgres:753159abc@172.17.0.1:5004/postgres?sslmode=disable")

	fmt.Println(pgHost)
	if err != nil {
		log.Fatal("sql could not open postgres database")
	}
	if Pg.Ping() != nil {
		log.Fatal("could not connect to postgres database")
	}
	r.Pg = Pg

}
