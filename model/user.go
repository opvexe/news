package model

import (
	"errors"
	"shumin-project/admin-blog-web/sqlinit"
)

//根据用户名查询用户
func Login(num string) (*User, error) {
	var user User
	err := sqlinit.Db.QueryRow(`select id,num,name,pass from user where num = ?`, num).Scan(&user.Id, &user.Num, &user.Num, &user.Pass)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//根据id查询用户
func GetUser(uid int64) (*User, error) {
	var user User
	err := sqlinit.Db.QueryRow(`select id,num,name,pass,phone,email,status,ctime from user where id = ? limit 1`, uid).Scan(&user.Id, &user.Num, &user.Name, &user.Pass, &user.Phone, &user.Email, &user.Status, &user.Ctime)
	if err != nil {
		return nil, err
	}
	return &user, err
}

//判断用户是否存在
func ExistsUser(num string) bool {
	var id int
	err := sqlinit.Db.QueryRow(`select * from user where num = ?`, num).Scan(&id)
	if err != nil {
		return false
	}
	return true
}

//添加用户
func AddUser(uer *User) error {
	tx, _ := sqlinit.Db.Begin()
	reslut, err := tx.Exec(`insert into user (num,name,pass,phone,email,status,ctime) values(?,?,?,?,?,?,?)`, uer.Num, uer.Pass, uer.Phone, uer.Email, uer.Status, uer.Ctime)
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

//修改用户
func AlterUser(user *User) error {
	tx, _ := sqlinit.Db.Begin()
	//开启事务
	result, err := tx.Exec(`update user set name = ?,phone = ?,email = ? where id = ?`, user.Name, user.Phone, user.Email, user.Id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		_ = tx.Rollback()
		return errors.New("事务失败")
	}
	return tx.Commit()
}

//删除用户
func DeleteUser(id int64) error {
	tx, _ := sqlinit.Db.Begin()
	result, err := tx.Exec(`delete from user  where id = ?`, id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		_ = tx.Rollback()
		return errors.New("事务失败")
	}
	return tx.Commit()
}

//获取用户总数
func CountUser() int {
	var count int
	err := sqlinit.Db.QueryRow(`select count(id) from user `).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

//分页
func Page(pi, ps int) ([]User, error) {
	q := make([]User, 0)
	rows, err := sqlinit.Db.Query(`select *from user limit ?,?`, (pi-1)*ps, ps)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Num, &user.Name, &user.Pass, &user.Phone, &user.Email, &user.Status, &user.Ctime)
		if err != nil {
			return nil, err
		}
		q = append(q, user)
	}
	return q, nil
}
