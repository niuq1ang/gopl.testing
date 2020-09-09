package db

import (
	"database/sql"
	"fmt"
	"runtime/debug"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ngaut/log"
	gorp "gopkg.in/gorp.v1"
)

var BangDBMap, NoticeDBMap *gorp.DbMap

func InitDB() (err error) {
	BangDBMap, err = OpenDB(config.Config.BangDBSpec)
	if err != nil {
		return
	}
	NoticeDBMap, err = OpenDB(config.Config.NoticeDBSpec)
	if err != nil {
		return
	}
	return
}

func OpenDB(dbSpec string) (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", dbSpec)
	if err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "utf8mb4"}}
	// dbmap.TraceOn("[SQL]", logger.New(os.Stdout, "[SQL]", log.Lmicroseconds))
	return dbmap, nil
}

func DBMTransact(src *gorp.DbMap, txFunc func(*gorp.Transaction) error) (err error) {
	tx, err := src.Begin()
	if err != nil {
		err = errors.Sql(err)
		return
	}
	defer func() {
		if p := recover(); p != nil {
			log.Warn("%s: %s", p, debug.Stack())
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
		}
		if err != nil {
			tx.Rollback()
			return
		}
		err = errors.Sql(tx.Commit())
	}()
	return txFunc(tx)
}
