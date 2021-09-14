package graph

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/goalsApp/server/db"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// Reads and parses the schema from file
// Associates root resolver, checks for errors along the way
func parseSchema(path string, resolver interface{}) *graphql.Schema {
	bstr, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("Couldn't get the Graphql schema file", err)
	}

	schemaString := string(bstr)
	parsedSchema, err := graphql.ParseSchema(
		schemaString,
		resolver,
	)

	if err != nil {
		log.Fatal("Couldn't parse the Graphql schema", err)
	}

	return parsedSchema
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		write.Header().Set("Access-Control-Allow-Origin", "*")
		write.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		if request.Method == "OPTIONS" {
			write.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(write, request)
	})
}

// Serves GraphQL Playground on root
// Serves GraphQL endpoint at /graphql
func ConnectGraphqQL(db db.DB) {
	playground := http.FileServer(http.Dir("graphql/graphqlPlayground"))

	http.Handle("/", playground)
	http.Handle("/graphql", CorsMiddleware(&relay.Handler{
		Schema: parseSchema("./graphql/schema.graphql", &RootResolver{db: db}),
	}))

	fmt.Println("serving on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
