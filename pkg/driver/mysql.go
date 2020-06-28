package driver

import "database/sql"

func Connect() (*sql.DB, error) {
	return sql.Open("mysql", "root:123321@/learn")
}