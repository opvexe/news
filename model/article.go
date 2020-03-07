package model

import (
	"errors"
	"shumin-project/admin-blog-web/sqlinit"
)

// 文章总数
func ArticleCount() int {
	var count int
	err := sqlinit.Db.QueryRow(`select count(id) from article`).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

// 分页
func ArticlePage(pi, ps int) ([]Article, error) {
	q := make([]Article, 0)
	rows, err := sqlinit.Db.Query(`select *from article order by id desc limit ?,?`, (pi-1)*ps, ps)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.Cid, &article.Uid, &article.Title, &article.Origin, &article.Author, &article.Content, &article.Hits, &article.Ctime, &article.Utime)
		if err != nil {
			return nil, err
		}
		q = append(q, article)
	}
	return q, nil
}

//删除
func ArticleDelete(id int64) error {
	tx, _ := sqlinit.Db.Begin()
	reslut, err := tx.Exec(`delete from article where id =?`, id)
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

//添加新闻
func ArticleAdd(art *Article) error {
	tx, _ := sqlinit.Db.Begin()
	reslut, err := tx.Exec(`insert into article (title,author,cid,content,hits,ctime,utime,origin,uid) values (?,?,?,?,?,?,?,?,?)`, art.Title, art.Author, art.Cid, art.Content, art.Hits, art.Ctime, art.Origin, art.Uid)
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
