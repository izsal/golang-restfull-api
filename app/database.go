package app

import (
	"database/sql"
	"time"

	"github.com/izsal/belajar-golang-restfull-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3333)/belajar_golang_database_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate create -ext sql -dir db/migrations create_table_second
	// migrate create -ext sql -dir db/migrations create_table_third
	// migrate create -ext sql -dir db/migrations sample_dirty_state
	// migrate -database "mysql://root:root@tcp(localhost:3333)/belajar_golang_database_migration" -path db/migrations up
	// migrate -database "mysql://root:root@tcp(localhost:3333)/belajar_golang_database_migration" -path db/migrations down
	// migrate -database "mysql://root:root@tcp(localhost:3333)/belajar_golang_database_migration" -path db/migrations version
	// migrate -database "mysql://root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations force 20240904024200 // This is a version before Dirty occurs to be able to improve the failed migrate
}
