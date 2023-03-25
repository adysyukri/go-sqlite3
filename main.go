package main

import (
	"context"
	"database/sql"
)

type account struct {
	id      int
	name    string
	balance float64
}

var db *sql.DB = new(sql.DB)

var ctx = context.Background()

const select1 = `
SELECT id, name, balance FROM accounts WHERE id = ?;
`

func main() {

}
