package main

import (
	"database/sql"
	"github.com/ahmedkhaeld/simplbank/api"
	db "github.com/ahmedkhaeld/simplbank/db/sqlc"
	"github.com/ahmedkhaeld/simplbank/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuration variables", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db")
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
