package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	// "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/yush/go-orders-graphql-api/graph"
	"github.com/yush/go-orders-graphql-api/graph/generated"
	"github.com/yush/go-orders-graphql-api/graph/model"
	"gorm.io/driver/postgres"
)

var db *gorm.DB;

func initDB() {
	dsn := "host=localhost user=postgres password=yushajay1 dbname=nooble port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        fmt.Println(err)
        panic("failed to connect database")
    }

    // db.Logger(true)

    // Create the database. This is a one-time step.
    // Comment out if running multiple times - You may see an error otherwise
    // db.Exec("CREATE DATABASE nooble")
    db.Exec("USE nooble")

    // Migration to create tables for Order and Item schema
    db.AutoMigrate(&model.Audio{})
	db.AutoMigrate(&model.Creator{})	
}



const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	initDB();
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
