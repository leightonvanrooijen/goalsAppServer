package main

import (
	"github.com/goalsApp/server/db"
	graph "github.com/goalsApp/server/graphql"
)

func main() {
	db := db.Connect()
	// db.Seed()
	graph.ConnectGraphqQL(db)
}
