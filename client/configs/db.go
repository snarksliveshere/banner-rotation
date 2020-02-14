package configs

import (
	"github.com/go-pg/pg"
	"sync"
)

var dbOnce sync.Once
var dbConn *pg.DB

type DB struct {
	Conf *AppConfig
}

func (db *DB) CreatePgConn() *pg.DB {
	opt := &pg.Options{
		Addr:     db.Conf.DBHost + ":" + db.Conf.DBPort,
		User:     db.Conf.DBUser,
		Password: db.Conf.DBPassword,
		Database: db.Conf.DBName,
	}

	dbOnce.Do(func() {
		dbConn = pg.Connect(opt)
	})
	return dbConn
}
