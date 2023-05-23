package main

import (
	"Projects/day/pkg/store/postgres"
	"fmt"
)

func main() {
	db, err := postgres.OpenDB("localhost", "", "postgres", "1234", "day")
	if err != nil {
		fmt.Printf("Fail connecting ")
	}
	defer db.Close()
}
