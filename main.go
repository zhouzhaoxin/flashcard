package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

func initRouter(r *gin.Engine) {
	// 记忆卡片
	r.GET("/", remember)
	r.GET("/remember/next", rememberNext)
	r.GET("/remember/prev", rememberPrev)
	r.GET("/known", known)

	// 卡片添加
	r.GET("/cards/add/index", cardsAddIndex)
	r.POST("/cards/add", cardsAddHandler)

	// 卡片删除
	r.GET("/cards/delete", cardsDeleteHandler)

	// 卡片编辑
	r.POST("/cards/edit", cardsEditHandler)
	r.GET("/cards/edit/index", cardsEditIndex)

	// 卡片列表
	r.GET("/cards", cards)
	r.GET("/cards/index", cardsIndex)

}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "flashcard:flashcard@tcp(127.0.0.1:3306)/flashcard?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err := db.Close()
		handleErr(err)
	}()
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	r.Use(cors.Default())
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	initRouter(r)
	log.Fatal(r.Run(":80")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
