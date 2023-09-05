package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Postgres *pgxpool.Pool

func main() {
	PsqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s application_name=%s", "localhost", "5432", "username", "password", "dbname", "require", "ServiceName")
	connect(PsqlInfo)
	Listen()
	defer Postgres.Close()
}

func connect(PsqlInfo string) {
	poolConfig, err := pgxpool.ParseConfig(PsqlInfo)
	if err != nil {
		fmt.Println("Error providing config:" + err.Error())
	}

	Postgres, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		fmt.Println("PostGres SQL Open failed with error -  " + err.Error())
	} else {
		fmt.Println("PostGres SQL Open sucessfull.")
	}
}

func Listen() error {
	ctx := context.Background()
	conn, err := Postgres.Acquire(ctx)
	if err != nil {
		fmt.Println("Unable to acquire connection:" + err.Error())
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, `LISTEN $1`, "demo")

	if err != nil {
		fmt.Println("Error occured while listening postgress notification:" + err.Error())
		return err
	}

	for {
		fmt.Println(":::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
		notification, err := conn.Conn().WaitForNotification(context.Background())
		if err != nil {
			fmt.Println("Error while waiting for notification:" + err.Error())
			return err
		}
		fmt.Println("Received notification:" + notification.Channel + " and payload: " + notification.Payload)
	}

}
