package sqlinit

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"strings"
)

var Db *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/news?charset=utf8mb4")
	if err != nil {
		log.Error("打开数据库失败:", err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error("数据库连接失败:", err)
		return
	}
	Db = db
	if err = createData(); err != nil {
		log.Error("创建表失败:", err)
		return
	}
}

//执行创建表语句
func createData() error {
	tablesqls := strings.Split(SQL_CREATETABLE, "/;")
	for _, s := range tablesqls {
		s = strings.Replace(strings.TrimSpace(s), "/;", ";", 1)
		if s == "" {
			continue
		}
		_, err := Db.Exec(s)
		if err != nil {
			log.Error("创建数据库表失败:", err)
			return err
		}
	}
	return nil
}

//初始化数据库数据
func initData() error {
	tx, err := Db.Begin()
	if err != nil {
		log.Error("加载数据库失败:", err)
		return err
	}
	tablesqls := strings.Split(SQL_DATA, ";")
	for _, s := range tablesqls {
		s = strings.TrimSpace(s)
		fmt.Println(s)
		if s == "" {
			continue
		}
		_, err := tx.Exec(s)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	return err
}
