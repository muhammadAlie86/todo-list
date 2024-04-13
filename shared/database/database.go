package database

import (
	"database/sql"
	"fmt"
	"time"
	"todo-list/shared/config"

	_ "github.com/lib/pq"
)

type Database struct {
	PGSQL *sql.DB
}
type IDatabase interface {
	GetDatabase() *sql.DB
	CloseDatabase()
}

func NewDataBase(cfg *config.AllConfig) IDatabase {
	PGSQL, err := initPostgreSQL(cfg.PGSQLConfig)

	if err != nil {
		panic(err)
	}

	return &Database{
		PGSQL: PGSQL,
	}
}

func initPostgreSQL(cfg *config.PGSQLConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.PG_HOST, cfg.PG_PORT, cfg.PG_USER, cfg.PG_DBNAME, cfg.PG_PASSWORD, cfg.PG_SSLMODE)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.PG_MaxIdleConn)
	db.SetMaxOpenConns(cfg.PG_MaxOpenConn)
	if cfg.PG_ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(cfg.PG_ConnMaxLifetime) * time.Minute)
	}
	if cfg.PG_ConnMaxIdleTime > 0 {
		db.SetConnMaxIdleTime(time.Duration(cfg.PG_ConnMaxIdleTime) * time.Minute)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

func (d *Database) GetDatabase() *sql.DB {
	return d.PGSQL
}
func (d *Database) CloseDatabase() {
	if err := d.PGSQL.Close(); err != nil {
		panic(err)
	}
}
