package pg_engine

import "database/sql"

type DBEngine interface {
	GetDB() *sql.DB
	Close()
}
