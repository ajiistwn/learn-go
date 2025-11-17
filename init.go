package main

import (
	"fmt"
	"learn-go/database"
	_ "learn-go/internal"
)

func main() {
	fmt.Println("Database Connection:", database.GetConnection())
}
