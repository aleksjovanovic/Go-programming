package main

import (
	// "context"
	"fmt"

	db "github.com/aleksjovanovic/go-rest-api/internal/database"
)

func Run() error {
	fmt.Println("Starting up our application ...")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	// if err := db.Ping(context.Background()); err != nil {
	// 	return err
	// }

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}
	fmt.Println("successfully connected and pinged database")
	return nil
}

func main() {
	fmt.Println("Go REST API ...")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
