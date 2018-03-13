package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {
	http.Handle("/", middleware(createSchema()))
	fmt.Println("Listen on :8080")
	http.ListenAndServe(":8080", nil)
}

func middleware(schema graphql.Schema) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryString, _ := ioutil.ReadAll(request.Body)
		fmt.Println("incoming query: ", string(queryString))

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(queryString),
		})

		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(result)
	}
}
