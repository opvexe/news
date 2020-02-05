package main

import (
	"shumin-project/admin-blog-web/router"
	_ "shumin-project/admin-blog-web/sqlinit"
)

func main() {
	router.Run()
}
