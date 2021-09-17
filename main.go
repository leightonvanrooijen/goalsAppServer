package main

import (
	"fmt"

	"github.com/goalsApp/server/db"
	graph "github.com/goalsApp/server/graphql"
)

func main() {
	fmt.Println("I ran!")
	db := db.Connect()
	db.Seed()
	graph.ConnectGraphqQL(db)
}
