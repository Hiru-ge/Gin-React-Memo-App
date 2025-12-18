package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Memo struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
}

func getAllMemos() ([]Memo, error) {
	var memos []Memo

	rows, err := db.Query("SELECT * FROM memos")
	if err != nil {
		return nil, fmt.Errorf("getAllMemos : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var memo Memo
		if err := rows.Scan(&memo.ID, &memo.Title, &memo.Content, &memo.Created_at); err != nil {
			return nil, fmt.Errorf("getAllMemos : %v", err)
		}
		memos = append(memos, memo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAllMemos : %v", err)
	}
	return memos, nil
}

func getMemoByID(id string) (Memo, error) {
	var memo Memo

	row := db.QueryRow("SELECT * FROM memos WHERE id = ?", id)
	if err := row.Scan(&memo.ID, &memo.Title, &memo.Content, &memo.Created_at); err != nil {
		if err == sql.ErrNoRows {
			return memo, fmt.Errorf("getAllMemos %d: %v", id, err)
		}
		return memo, fmt.Errorf("getAllMemos %d: %v", id, err)
	}
	return memo, nil
}

func getAllMemosHandler(c *gin.Context) {
	memos, err := getAllMemos()
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, memos)
}

func getMemoByIDHandler(c *gin.Context) {
	id := c.Param("id")
	memo, err := getMemoByID(id)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, memo)
}

func main() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "memo_app"
	cfg.ParseTime = true // DB内部では[]byteで扱われているcreated_atを正しくtime.Timeで解釈するための設定

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// DB接続確認
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// ルーティング
	router := gin.Default()
	router.GET("/memos", getAllMemosHandler)
	router.GET("/memos/:id", getMemoByIDHandler)

	router.Run("localhost:8080")
}
