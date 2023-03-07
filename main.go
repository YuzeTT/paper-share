package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	FileList  string `json:"file_list"`
	Author    string `json:"author"`
	AuthorIP  string `json:"author_ip"`
	CreatedTS string `json:"created_ts"`
	UpdatedTS string `json:"updated_ts"`
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func main() {
	log.Println("[-] 程序开始初始化...")
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// var version string
	// err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 创建表
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS pages (id TEXT NOT NULL UNIQUE,title TEXT,content TEXT,file_list TEXT,author BLOB,author_ip INTEGER NOT NULL,created_ts INTEGER NOT NULL,updated_ts INTEGER NOT NULL,PRIMARY KEY(id) )")
	if err != nil {
		log.Println("[x] 创建数据表出错")
	} else {
		log.Println("[√] 数据表初始化成功")
	}
	statement.Exec()

	log.Println("[-] 开始启动HTTP服务...")

	router := gin.Default()
	// gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)
	router.Use(Cors())

	api := router.Group("/api")

	api.GET("/pages", func(ctx *gin.Context) {
		rows, err := db.Query("SELECT * FROM pages")
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		pages := []Page{}
		for rows.Next() {
			var page Page
			if err := rows.Scan(&page.ID, &page.Title, &page.Content, &page.FileList, &page.Author, &page.AuthorIP, &page.CreatedTS, &page.UpdatedTS); err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			pages = append(pages, page)
		}

		ctx.JSON(http.StatusOK, pages)
	})

	log.Println("[√] HTTP服务已启动")
	router.Run("127.0.0.1:3000")
}
