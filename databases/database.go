package databases

import (
	"github.com/jmoiron/sqlx"
)

type Database interface {
	Connect() *sqlx.DB
}
