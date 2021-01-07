package graphql

import (
	gql "github.com/mattdamon108/gqlmerge/lib"
	"log"
)

func NewSchema() *string {
	schema := gql.Merge("  ", "./delivery/graphql/schema")

	log.Println(schema)

	return schema
}
