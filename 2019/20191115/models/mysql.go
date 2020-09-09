package models

import (
	"database/sql"
	"github.com/bangwork/mygo/20191115/config"
	_ "github.com/go-sql-driver/mysql" // loag mysql driver
	gorp "gopkg.in/gorp.v1"
)

var ProjectDBMap, WikiDBMap *gorp.DbMap

func InitDB() error {
	var err error
	ProjectDBMap, err = OpenDB(config.Config.ProjectDBSpec)
	if err != nil {
		return err
	}
	WikiDBMap, err = OpenDB(config.Config.WikiDBSpec)
	if err != nil {
		return err
	}
	return nil
}

func OpenDB(dbSpec string) (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", dbSpec)
	if err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "utf8mb4"}}
	// dbmap.TraceOn("[gorp]", log.New(os.Stdout, "Sql:", log.Lmicroseconds))
	dbmap.TraceOff()
	return dbmap, nil
}
