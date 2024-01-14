package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	fileName := "a.csv"
	databaseUrl := "postgres://postgres:password@localhost:5432/postgres"

	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	start := time.Now()

	res, err := conn.PgConn().CopyFrom(context.Background(), file, "COPY \"slug-odin\" FROM STDIN (FORMAT csv)")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.RowsAffected())

	elapsed := time.Since(start)
	fmt.Println("it took", elapsed)
}
