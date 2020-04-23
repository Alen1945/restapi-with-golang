package postgres

import (
	"github.com/go-pg/pg/v9"
)

func New(options *pg.Options) *pg.DB {
	db := pg.Connect(options)
	return db
}
