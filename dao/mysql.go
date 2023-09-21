/*
Package dao is the Dao layer of the program.
This program uses two databases, BoltDB and MySQL.
If you want to keep it simple, starting Simple Mode will default to the embedded
database BoltDB.
*/

package dao

import "database/sql"

type MySQL struct {
	db           *sql.DB
	Username     string
	password     string
	url          string
	port         string
	databaseName string
}
