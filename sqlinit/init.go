package sqlinit

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var Db *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/news?charset=utf8mb4")
	if err != nil {
		fmt.Println("打开数据库失败:", err)
		return
	}
	if err = db.Ping(); err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	Db = db
	if err = createData(); err != nil {
		fmt.Println("创建表失败:", err)
		return
	}
	//if err = initData(); err != nil {
	//	fmt.Println("初始化数据失败:", err)
	//	return
	//}
}

//执行创建表语句
func createData() error {
	tablesqls := strings.Split(SQL_CREATETABLE, "/;")
	for _, sql := range tablesqls {
		sql = strings.Replace(strings.TrimSpace(sql), "/;", ";", 1)
		if sql == "" {
			continue
		}
		_, err := Db.Exec(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

//初始化数据库数据
func initData() error {
	tx, err := Db.Begin()
	tablesqls := strings.Split(SQL_DATA, ";")
	for _, sql := range tablesqls {
		sql = strings.TrimSpace(sql)
		fmt.Println(sql)
		if sql == "" {
			continue
		}
		_, err := tx.Exec(sql)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	return err
}
