package model

import (
	"errors"
	"shumin-project/admin-blog-web/sqlinit"
)

//分页
func ClassPage(pi, ps int) ([]Class, error) {
	q := make([]Class, 0)
	rows, err := sqlinit.Db.Query(`select *from class limit ?,?`, (pi-1)*ps, ps)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var class Class
		err := rows.Scan(&class.Id, &class.Name, &class.Description)
		if err != nil {
			return nil, err
		}
		q = append(q, class)
	}
	return q, nil
}

//获取总数
func ClassCount() int {
	var count int
	err := sqlinit.Db.QueryRow(`select count(id) from class`).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

//获取所有
func ClassAll() ([]Class, error) {
	q := make([]Class, 0)
	rows, err := sqlinit.Db.Query(`select * from class`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var class Class
		err := rows.Scan(&class.Id, &class.Name, &class.Description)
		if err != nil {
			return nil, err
		}
		q = append(q, class)
	}
	return q, nil
}

//查询某一条
func ClassGet(id int64) (*Class, error) {
	var class Class
	err := sqlinit.Db.QueryRow(`select * from class where id =? limit 1`, id).Scan(&class.Id, &class.Name, &class.Description)
	if err != nil {
		return nil, err
	}
	return &class, nil
}

// Class 添加
func ClassAdd(class *Class) error {
	tx, _ := sqlinit.Db.Begin()
	reslut, err := tx.Exec(`insert into class(name,description) values(?,?)`, class.Name, class.Description)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	rows, _ := reslut.RowsAffected()
	if rows < 1 {
		_ = tx.Rollback()
		return errors.New("事务失败")
	}
	return tx.Commit()
}

//修改
func ClassEdit(class *Class) error {
	tx, _ := sqlinit.Db.Begin()
	reslut, err := tx.Exec(`update class set name =?,description = ? where id =?`, class.Name, class.Description, class.Id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	rows, _ := reslut.RowsAffected()
	if rows < 1 {
		_ = tx.Rollback()
		return errors.New("事务失败")
	}
	return tx.Commit()
}

// 删除
func ClassDelete(id int64) error {
	tx, _ := sqlinit.Db.Begin()
	reslut, err := tx.Exec(`delete from class where id =?`, id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	rows, _ := reslut.RowsAffected()
	if rows < 1 {
		_ = tx.Rollback()
		return errors.New("事务失败")
	}
	return tx.Commit()
}

//查询
func ClassNameById(id int64) string {
	var name string
	err := sqlinit.Db.QueryRow(`select name from class where id =?`, id).Scan(&name)
	if err != nil {
		return ""
	}
	return name
}

//查询
func ClassNameByIds(ids []int64) map[int64]string {
	dict := make(map[int64]string)
	rows, err := sqlinit.Db.Query(`select * from class where id in(?)`, ids)
	if err != nil {
		return nil
	}
	for rows.Next() {
		var class Class
		err := rows.Scan(&class.Id, &class.Name)
		if err != nil {
			return nil
		}
		dict[int64(class.Id)] = class.Name
	}
	return dict
}
