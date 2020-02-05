package model

import (
	"github.com/dgrijalva/jwt-go"
)

// JWT
type UserClaims struct {
	Id   int    `json:"id"`
	Num  string `json:"num"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// User
type User struct {
	Id     int    `json:"id"`
	Num    string `json:"num"`
	Name   string `json:"name"`
	Pass   string `json:"pass"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Status int    `json:"status"`
	Ctime  string `json:"ctime"`
}

// Class
type Class struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Article
type Article struct {
	Id      int64  `json:"id"`
	Cid     int64  `json:"cid"`
	Uid     int64  `json:"uid"`
	Title   string `json:"title"`
	Origin  string `json:"origin"`
	Author  string `json:"author"`
	Content string `json:"content"`
	Hits    int64  `json:"hits"`
	Utime   string `json:"utime"`
	Ctime   string `json:"ctime"`
	Status  bool   `json:"status"`
}
