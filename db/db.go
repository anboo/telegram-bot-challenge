package db

import (
	"context"
	"database/sql"
	"log"
	"sync"
)

type Database struct {
	dsn string
	db  *sql.DB
	o   sync.Once
}

func CreateDatabase(dsn string) *Database {
	return &Database{dsn: dsn, o: sync.Once{}}
}

func (d *Database) Conn(ctx context.Context) *sql.DB {
	if d.db == nil {
		d.connect()
	} else {
		err := d.db.PingContext(ctx)
		if err != nil {
			log.Println("database ping connection error " + err.Error())
			d.connect()
		}
	}

	return d.db
}

func (d *Database) connect() {
	d.o.Do(func() {
		var err error
		d.db, err = sql.Open("postgres", d.dsn)
		if err != nil {
			panic("error database connection " + err.Error())
		}
	})
}
