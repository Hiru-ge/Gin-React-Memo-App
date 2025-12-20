package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Memo struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
}

func getAllMemos() ([]Memo, error) {
	var memos []Memo

	rows, err := db.Query("SELECT id, title, content, created_at FROM memos")
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

func getMemoByID(id int64) (Memo, error) {
	var memo Memo

	row := db.QueryRow("SELECT id, title, content, created_at FROM memos WHERE id = ?", id)
	if err := row.Scan(&memo.ID, &memo.Title, &memo.Content, &memo.Created_at); err != nil {
		if err == sql.ErrNoRows {
			return memo, fmt.Errorf("getMemoByID %d: %v", id, err)
		}
		return memo, fmt.Errorf("getMemoByID %d: %v", id, err)
	}
	return memo, nil
}

func addMemo(memo Memo) (Memo, error) {
	result, err := db.Exec("INSERT INTO memos (title, content) VALUES (?, ?)", memo.Title, memo.Content)
	if err != nil {
		return Memo{}, fmt.Errorf("addMemo: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Memo{}, fmt.Errorf("addMemo: %v", err)
	}
	newMemo, err := getMemoByID(id)
	if err != nil {
		return Memo{}, fmt.Errorf("addMemo: %v", err)
	}
	return newMemo, nil
}

func editMemo(memo Memo) (Memo, error) {
	result, err := db.Exec("UPDATE memos SET title = ?, content = ? WHERE id = ?", memo.Title, memo.Content, memo.ID)
	if err != nil {
		return Memo{}, fmt.Errorf("editMemo: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return Memo{}, fmt.Errorf("editMemo: %v", err)
	}
	if rows == 0 {
		// 変更行が0でも「変更前と同じ値で更新されたケース」があり得るので存在確認し、存在しない場合だけエラーを返す
		if _, err := getMemoByID(memo.ID); err == sql.ErrNoRows {
			return Memo{}, sql.ErrNoRows
		}
	}
	editedMemo, err := getMemoByID(memo.ID)
	if err != nil {
		return Memo{}, fmt.Errorf("editMemo: %v", err)
	}
	return editedMemo, nil
}

func deleteMemoByID(id int64) error {
	result, err := db.Exec("DELETE FROM memos WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("deleteMemoByID: %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("deleteMemoByID: %v", err)
	}
	if rows == 0 {
		// 削除対象の行が最初から存在しない際にエラーを返す
		return sql.ErrNoRows
	}
	return nil
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
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	memo, err := getMemoByID(id)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, memo)
}

func addMemoHandler(c *gin.Context) {
	var newMemo Memo
	if err := c.BindJSON(&newMemo); err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	newMemo, err := addMemo(newMemo)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newMemo)
}

func editMemoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	var editedMemo Memo
	if err := c.BindJSON(&editedMemo); err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	editedMemo.ID = id
	editedMemo, err = editMemo(editedMemo)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, editedMemo)
}

func deleteMemoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	err = deleteMemoByID(id)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
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
	router.POST("/memos", addMemoHandler)
	router.PUT("/memos/:id", editMemoHandler)
	router.DELETE("/memos/:id", deleteMemoHandler)

	router.Run("localhost:8080")
}
