package main

import (
	_ "backend-API/docs"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// エラーレスポンスの形を定義します
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

func getMemos(db *sql.DB) ([]Memo, error) {
	var memos []Memo

	rows, err := db.Query("SELECT id, title, content, created_at FROM memos")
	if err != nil {
		return nil, fmt.Errorf("getMemos: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var memo Memo
		if err := rows.Scan(&memo.ID, &memo.Title, &memo.Content, &memo.CreatedAt); err != nil {
			return nil, fmt.Errorf("getMemos: %v", err)
		}
		memos = append(memos, memo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getMemos: %v", err)
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

func createMemo(db *sql.DB, memo Memo) (Memo, error) {
	result, err := db.Exec("INSERT INTO memos (title, content) VALUES (?, ?)", memo.Title, memo.Content)
	if err != nil {
		return Memo{}, fmt.Errorf("createMemo: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Memo{}, fmt.Errorf("createMemo: %v", err)
	}
	newMemo, err := getMemoByID(db, id)
	if err != nil {
		return Memo{}, fmt.Errorf("createMemo: %v", err)
	}
	return newMemo, nil
}

func updateMemo(db *sql.DB, memo Memo) (Memo, error) {
	result, err := db.Exec("UPDATE memos SET title = ?, content = ? WHERE id = ?", memo.Title, memo.Content, memo.ID)
	if err != nil {
		return Memo{}, fmt.Errorf("updateMemo: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return Memo{}, fmt.Errorf("updateMemo: %v", err)
	}
	if rows == 0 {
		// 変更行が0でも「変更前と同じ値で更新されたケース」があり得るので存在確認し、存在しない場合だけエラーを返す
		if _, err := getMemoByID(db, memo.ID); err == sql.ErrNoRows {
			return Memo{}, sql.ErrNoRows
		}
	}
	updatedMemo, err := getMemoByID(db, memo.ID)
	if err != nil {
		return Memo{}, fmt.Errorf("updateMemo: %v", err)
	}
	return updatedMemo, nil
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

// getMemosHandler godoc
// @Summary      メモ一覧取得
// @Description  全てのメモを取得します
// @Tags         memos
// @Accept       json
// @Produce      json
// @Success      200  {array}  Memo
// @Router       /memos [get]
func (s *Server) getMemosHandler(c *gin.Context) {
	memos, err := getMemos(s.db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, memos)
}

// getMemoByIDHandler godoc
// @Summary      メモ詳細取得
// @Description  IDを指定して特定のメモを取得します
// @Tags         memos
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Memo ID"
// @Success      200  {object}  Memo
// @Failure      400  {object}  HTTPError  "Invalid ID format"
// @Failure      404  {object}  HTTPError  "Memo not found"
// @Router       /memos/{id} [get]
func (s *Server) getMemoByIDHandler(c *gin.Context) {
	id, ok := s.parseID(c)
	if !ok {
		return
	}
	memo, err := getMemoByID(s.db, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Memo not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, memo)
}

// createMemoHandler godoc
// @Summary      メモ新規作成
// @Description  新しいメモを作成します
// @Tags         memos
// @Accept       json
// @Produce      json
// @Param        request body   Memo  true  "Memo content"
// @Success      201     {object}  Memo
// @Failure      400     {object}  HTTPError  "Invalid input"
// @Failure      500     {object}  HTTPError  "Server error"
// @Router       /memos [post]
func (s *Server) createMemoHandler(c *gin.Context) {
	var newMemo Memo
	if err := c.BindJSON(&newMemo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newMemo, err := createMemo(s.db, newMemo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newMemo)
}

// updateMemoHandler godoc
// @Summary      メモ編集
// @Description  IDを指定してメモの内容を更新します
// @Tags         memos
// @Accept       json
// @Produce      json
// @Param        id      path   int   true  "Memo ID"
// @Param        request body   Memo  true  "Updated content"
// @Success      200     {object}  Memo
// @Failure      400     {object}  HTTPError  "Invalid input"
// @Failure      500     {object}  HTTPError  "Server error"
// @Router       /memos/{id} [put]
func (s *Server) updateMemoHandler(c *gin.Context) {
	id, ok := s.parseID(c)
	if !ok {
		return
	}
	var editedMemo Memo
	if err := c.BindJSON(&editedMemo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	editedMemo.ID = id
	editedMemo, err := updateMemo(s.db, editedMemo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, editedMemo)
}

// deleteMemoHandler godoc
// @Summary      メモ削除
// @Description  IDを指定してメモを削除します
// @Tags         memos
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Memo ID"
// @Success      204  "No Content"
// @Failure      400  {object}  HTTPError  "Invalid ID format"
// @Failure      500  {object}  HTTPError  "Server error"
// @Router       /memos/{id} [delete]
func (s *Server) deleteMemoHandler(c *gin.Context) {
	id, ok := s.parseID(c)
	if !ok {
		return
	}
	err := deleteMemoByID(s.db, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (s *Server) parseID(c *gin.Context) (int64, bool) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return 0, false
	}
	return id, true
}

// @title           Memo App API
// @version         1.0
// @description     Ginで作られたメモアプリのAPIサーバーです
// @host            localhost:8080
// @BasePath        /
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
	router.GET("/memos", server.getMemosHandler)
	router.GET("/memos/:id", server.getMemoByIDHandler)
	router.POST("/memos", server.createMemoHandler)
	router.PUT("/memos/:id", server.updateMemoHandler)
	router.DELETE("/memos/:id", server.deleteMemoHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("localhost:8080")
}
