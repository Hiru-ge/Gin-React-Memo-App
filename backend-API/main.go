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

type Server struct {
	db *sql.DB
}

type Memo struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func getAllMemos(db *sql.DB) ([]Memo, error) {
	var memos []Memo

	rows, err := db.Query("SELECT id, title, content, created_at FROM memos")
	if err != nil {
		return nil, fmt.Errorf("getAllMemos : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var memo Memo
		if err := rows.Scan(&memo.ID, &memo.Title, &memo.Content, &memo.CreatedAt); err != nil {
			return nil, fmt.Errorf("getAllMemos : %v", err)
		}
		memos = append(memos, memo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAllMemos : %v", err)
	}
	return memos, nil
}

func getMemoByID(db *sql.DB, id int64) (Memo, error) {
	var memo Memo

	row := db.QueryRow("SELECT id, title, content, created_at FROM memos WHERE id = ?", id)
	if err := row.Scan(&memo.ID, &memo.Title, &memo.Content, &memo.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return memo, fmt.Errorf("getMemoByID %d: %v", id, err)
		}
		return memo, fmt.Errorf("getMemoByID %d: %v", id, err)
	}
	return memo, nil
}

func addMemo(db *sql.DB, memo Memo) (Memo, error) {
	result, err := db.Exec("INSERT INTO memos (title, content) VALUES (?, ?)", memo.Title, memo.Content)
	if err != nil {
		return Memo{}, fmt.Errorf("addMemo: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Memo{}, fmt.Errorf("addMemo: %v", err)
	}
	newMemo, err := getMemoByID(db, id)
	if err != nil {
		return Memo{}, fmt.Errorf("addMemo: %v", err)
	}
	return newMemo, nil
}

func editMemo(db *sql.DB, memo Memo) (Memo, error) {
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
		if _, err := getMemoByID(db, memo.ID); err == sql.ErrNoRows {
			return Memo{}, sql.ErrNoRows
		}
	}
	editedMemo, err := getMemoByID(db, memo.ID)
	if err != nil {
		return Memo{}, fmt.Errorf("editMemo: %v", err)
	}
	return editedMemo, nil
}

func deleteMemoByID(db *sql.DB, id int64) error {
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

func (s *Server) getAllMemosHandler(c *gin.Context) {
	memos, err := getAllMemos(s.db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, memos)
}

func (s *Server) getMemoByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	memo, err := getMemoByID(s.db, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Memo not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, memo)
}

func (s *Server) addMemoHandler(c *gin.Context) {
	var newMemo Memo
	if err := c.BindJSON(&newMemo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newMemo, err := addMemo(s.db, newMemo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newMemo)
}

func (s *Server) editMemoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var editedMemo Memo
	if err := c.BindJSON(&editedMemo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	editedMemo.ID = id
	editedMemo, err = editMemo(s.db, editedMemo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, editedMemo)
}

func (s *Server) deleteMemoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = deleteMemoByID(s.db, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	db, err := sql.Open("mysql", cfg.FormatDSN())
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
	server := &Server{db: db}
	router := gin.Default()
	router.GET("/memos", server.getAllMemosHandler)
	router.GET("/memos/:id", server.getMemoByIDHandler)
	router.POST("/memos", server.addMemoHandler)
	router.PUT("/memos/:id", server.editMemoHandler)
	router.DELETE("/memos/:id", server.deleteMemoHandler)
	router.Run("localhost:8080")
}
