# Gin + React Memo App

Go (Gin) と React (TypeScript) を用いた、SPA構成のシンプルなメモアプリケーションです。
RESTful API の設計と実装、およびフロントエンドとの非同期通信の学習を目的としています。

## 技術スタック

### Backend
- **Language**: Go
- **Framework**: Gin Web Framework
- **Database Driver**: database/sql + go-sql-driver/mysql
- **Database**: MySQL 8.0

### Frontend
- **Library**: React
- **Language**: TypeScript
- **Build Tool**: Vite
- **HTTP Client**: axios (または fetch)

---

## データベース設計 (Schema)

データベース名: `memo_app`

### `memos` テーブル
メモの本体を保存するテーブルです。

| Column Name | Type | Key | Note |
| :--- | :--- | :--- | :--- |
| `id` | BIGINT | PK | AUTO_INCREMENT |
| `title` | VARCHAR(255) | | メモのタイトル (Not Null) |
| `content` | TEXT | | メモの本文 |
| `created_at` | DATETIME | | 作成日時 (Default: CURRENT_TIMESTAMP) |

```sql
-- 初期化用SQL
CREATE TABLE memos (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```