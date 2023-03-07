package main

import (
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	FileList  string `json:"file_list"`
	Author    string `json:"author"`
	AuthorIP  string `json:"author_ip"`
	CreatedTS int    `json:"created_ts"`
	UpdatedTS int    `json:"updated_ts"`
}

type Post struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

//go:embed web/dist
var dist embed.FS

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

	distFS, _ := fs.Sub(dist, "dist")
	router.StaticFS("/static", http.FS(distFS))

	router.NoRoute(func(c *gin.Context) {
		c.FileFromFS("index.html", http.FS(dist))
	})
	api.GET("/pages", func(ctx *gin.Context) {
		rows, err := db.Query("SELECT * FROM pages order by created_ts desc")
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

	api.GET("/page", func(ctx *gin.Context) {
		id := ctx.Query("id")
		var page Page

		err := db.QueryRow("SELECT * FROM pages WHERE id = ?", id).Scan(&page.ID, &page.Title, &page.Content, &page.FileList, &page.Author, &page.AuthorIP, &page.CreatedTS, &page.UpdatedTS)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, page)
	})

	api.POST("/post", func(ctx *gin.Context) {
		var post Post
		if err := ctx.ShouldBindJSON(&post); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		var sql_id string
		sql_id, err := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-", 8)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		result, err := db.Exec(`INSERT INTO pages
		(id, title, content, file_list, author, author_ip, created_ts, updated_ts)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, sql_id, post.Title, post.Content, "{}", post.Author, ctx.ClientIP(),
			time.Now().UnixNano()/1e6,
			time.Now().UnixNano()/1e6)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		id, _ := result.LastInsertId()
		log.Println(id)

		ctx.JSON(http.StatusCreated, post)
	})

	log.Println("[√] HTTP服务已启动")
	router.Run("localhost:3000")
}
