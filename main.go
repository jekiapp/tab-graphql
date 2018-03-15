package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {
	http.Handle("/graphql", middleware(createSchema()))
	fmt.Println("Listen on :8080")
	log.Println(http.ListenAndServe(":8080", nil))
}

func middleware(schema graphql.Schema) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryString, _ := ioutil.ReadAll(request.Body)

		// to handle cors by voyager
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		if request.Method == http.MethodOptions {
			return
		}

		// to cover voyager request
		// voyager will send request with structure {query:"{realquery}"}
		q := struct {
			Query string `json:"query"`
		}{}
		if err := json.Unmarshal(queryString, &q); err == nil {
			queryString = []byte(q.Query)
		}

		fmt.Println("incoming query: ", string(queryString))

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(queryString),
		})

		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(result)
	}
}
