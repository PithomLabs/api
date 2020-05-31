package graphql

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/graph-gophers/graphql-go"
)

//Schema respresents graphql schema
//Schema is parsed from internal/initialize/main.go
var Schema *graphql.Schema

//ExecuteQuery is a wrapper on schema.Exec() method
func ExecuteQuery(req *http.Request) ([]byte, error) {
	query := req.URL.Query().Get("query")
	resp := Schema.Exec(context.Background(), query, "", nil)
	json, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		return []byte{}, err
	}
	return json, nil
}
