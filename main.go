package main

import (
        "github.com/graphql-go/graphql"
        "github.com/graphql-go/handler"
        "log"
        "net/http"
)

func main() {
        // Schema
        fields := graphql.Fields{
                "hello": &graphql.Field{
                        Type: graphql.String,
                        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                                return "world", nil
                        },
                },
                "data": &graphql.Field{
                        Type: graphql.Boolean,
                        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                                return p.Info.FieldName, nil
                        },
                },
        }
        rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
        schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
        schema, err := graphql.NewSchema(schemaConfig)
        if err != nil {
                log.Fatalf("failed to create new schema, error: %v", err)
        }

        h := handler.New(&handler.Config{
                Schema:   &schema,
                Pretty:   true,
                GraphiQL: true,
        })

        http.Handle("/graphql", h)
        err = http.ListenAndServe(":9000", nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}
