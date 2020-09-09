package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	gorp "gopkg.in/gorp.v1"
	"log"
	"os"
	"regexp"
	"runtime/debug"
	"strings"
	"time"
)

func main() {
	fmt.Println("lalalaxxxxxxxx")
	MySQLMigrationFlow("", migrate)
}

const (
	DSNPattern = `%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FShanghai&charset=utf8mb4`
)

func migrate(tx *gorp.Transaction) error {
	return nil
}

type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	Verbose  bool
}

func (c *MySQLConfig) BindFlag() {
	flag.StringVar(&c.Host, "host", "127.0.0.1", "MySQL host")
	flag.IntVar(&c.Port, "port", 3306, "MySQL port")
	flag.StringVar(&c.User, "user", "root", "MySQL user")
	flag.StringVar(&c.Password, "password", "123456", "MySQL password")
	flag.StringVar(&c.DBName, "db", "bang", "MySQL database name")
	flag.BoolVar(&c.Verbose, "verbose", true, "Enable verbose log")
}

type MySQLDB struct {
	DB  *sql.DB
	DBM *gorp.DbMap
}

func InitMySQL(c *MySQLConfig) (*MySQLDB, error) {
	dsn := fmt.Sprintf(DSNPattern, c.User, c.Password, c.Host, c.Port, c.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	dbm := &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "utf8mb4",
		},
	}
	if c.Verbose {
		dbm.TraceOn("[SQL]", log.New(os.Stdout, "", log.Lmicroseconds))
	} else {
		dbm.TraceOff()
	}
	result := &MySQLDB{
		DB:  db,
		DBM: dbm,
	}
	return result, nil
}

func MySQLTransact(dbm *gorp.DbMap, txFunc func(tx *gorp.Transaction) error) (err error) {
	tx, err := dbm.Begin()
	if err != nil {
		return
	}
	time.Sleep(time.Minute * 5)

	defer func() {
		if p := recover(); p != nil {
			log.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			log.Printf("%s: %s\n", p, debug.Stack())
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
		err = tx.Commit()
	}()
	return txFunc(tx)
}

func MySQLMigrationFlow(alter string, migrate func(tx *gorp.Transaction) error) {
	c := new(MySQLConfig)
	c.BindFlag()
	flag.Parse()
	sqldb, err := InitMySQL(c)
	if err != nil {
		log.Fatalln(err)
	}

	if len(alter) > 0 {
		err = MySQLTransact(sqldb.DBM, func(tx *gorp.Transaction) error {
			return MySQLExecMultipleStatements(tx, alter)
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
	err = MySQLTransact(sqldb.DBM, migrate)
	if err != nil {
		log.Fatalln(err)
	}
}

func MySQLExecMultipleStatements(tx *gorp.Transaction, rawquery string) error {
	// 处理 SQL 字符串中带有分号的特殊情况
	re := regexp.MustCompile(`[^\\];`)
	locs := re.FindAllStringIndex(rawquery, -1)
	prev := 0
	for _, loc := range locs {
		end := loc[1]
		query := strings.TrimSpace(rawquery[prev:end])
		if query == "" || query == ";" {
			continue
		}
		if _, err := tx.Exec(query); err != nil {
			return err
		}
		prev = end
	}
	return nil
}
