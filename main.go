package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

func initRouter(r *gin.Engine) {
	r.GET("/ping", pong)
	r.GET("/", remember)
	r.GET("/remember/next", rememberNext)
	r.GET("/remember/prev", rememberPrev)
	r.GET("/cards", cards)
	r.GET("/add/code/index", addCodeIndex)
	r.GET("/add/vocab/index", addVocabIndex)
	r.POST("/add/vocab", addVocab)
	r.POST("/add/code", addCode)
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/flashcard?parseTime=true")
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
	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
