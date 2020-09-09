package db

import (
	"database/sql"
	"fmt"
	"github.com/bangwork/mygo/20190530/bi-week/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"

	gorp "gopkg.in/gorp.v1"
)

var PrivateDBMap *gorp.DbMap

func InitPrivateDB() error {
	var err error
	PrivateDBMap, err = OpenDB(
		config.StringOrPanic("private_db_user"),
		config.StringOrPanic("private_db_password"),
		config.StringOrPanic("private_db_ip"),
		config.StringOrPanic("private_db_name"),
	)
	return err
}

func OpenDB(user, password, ip, name string) (*gorp.DbMap, error) {
	dataSourceName := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&loc=Asia%sShanghai&parseTime=true`, user, password, ip, name, `%2F`)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "utf8mb4"}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "Sql:", log.Lmicroseconds))
	dbmap.TraceOff()
	return dbmap, nil
}
