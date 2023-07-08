package main

import (
	"github.com/amar-jay/go-api-boilerplate/cmd"
)

func main() {
	_, db, _, _ := cmd.Initialize()
	db = db.Debug()
	cmd.Migrate(db)
}
