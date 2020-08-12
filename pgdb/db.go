package pgdb

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	DB     *sql.DB
	Host   string
	Port   int
	DbName string
}

func NewPostgres() *Postgres {
	return &Postgres{
		Host:   "localhost",
		Port:   5432,
		DbName: "cekicilis",
	}
}

func (p *Postgres) Connect() error {
	var err error
	p.DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable", p.Host, p.Port, p.DbName))
	if err != nil {
		return err
	}

	if err := p.DB.Ping(); err != nil {
		return err
	}

	return nil
}
