package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	db "github.com/simplebank/db/sqlc"
)

func main() {
	fmt.Println("Running this")
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://root:secret@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	// list all authors
	authors, err := queries.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    "John Doe",
		Balance:  1000,
		Currency: "USD",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(authors)
}
