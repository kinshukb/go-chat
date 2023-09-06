package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func main() {
	var err error
	pool, err = pgxpool.New(context.Background(), "your-pg-host")

	// TODO: try to avoid using env vars, look for another api
	encryptedPwd := os.Getenv("postgresqlpwd")
	nonce := os.Getenv("postgresqlno")
	// cleartextPwd, err := encryptUtil.Decrypt(encryptedPwd, nonce)
	// if err != nil {
	// 	fmt.Println("Decrypt failed with error - " + err.Error())
	// }
	os.Setenv("PGPASSWORD", cleartextPwd)
	fmt.Println(cleartextPwd)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to connect to database:", err)
		os.Exit(1)
	}

	fmt.Println(`Type a message and press enter.
This message should appear in any other chat instances connected to the same
database.
Type "exit" to quit.`)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "exit" {
			os.Exit(0)
		}

		_, err = pool.Exec(context.Background(), "select pg_notify('chat', $1)", msg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error sending notification:", err)
			os.Exit(1)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error scanning from stdin:", err)
		os.Exit(1)
	}
}
