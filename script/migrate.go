package main

import (
	"github.com/amar-jay/go-api-boilerplate/cmd"
)

func main() {
	_, db, _, _ := cmd.Initialize()
	cmd.Migrate(db)
}
