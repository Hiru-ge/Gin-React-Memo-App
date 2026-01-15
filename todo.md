# todo.md

## 概要
シンプルなメモアプリ（Gin + React）の開発タスクを、ざっくりとまとめたチェックリストです。

---

## DB準備
- [x] MySQL にデータベース `memo_app` を作成する
  - 例: `CREATE DATABASE memo_app CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;`
- [x] `memos` テーブルを作成する
  - 例:
    ```sql
    CREATE TABLE memos (
      id BIGINT AUTO_INCREMENT PRIMARY KEY,
      title VARCHAR(255) NOT NULL,
      content TEXT,
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    ```
- [x] （任意）サンプルデータを投入する

---

## Backend (Go / Gin)
- [x] `main.go` を作成して Gin サーバを起動する
- [x] `database/sql` と `go-sql-driver/mysql` で DB 接続を実装（DSN は環境変数で管理）
- [x] モデルと CRUD を実装（`GET /memos`, `GET /memos/:id`, `POST /memos`, `PUT /memos/:id`, `DELETE /memos/:id`）
- [x] Postman / curl で API を手動確認する
- [x] リファクタリング

---

## Frontend (React + TypeScript)
- [x] Vite で React + ReactRouterプロジェクト初期化

### 1. メモ一覧ページ (`/`)
- [x] ルーティング設定（`/`）
- [x] モックデータで一覧ページを作成
- [ ] API 呼び出し関数（`GET /memos`）を作成
- [ ] API 連携して実際のデータを表示

### 2. メモ詳細ページ (`/memos/:id`)
- [ ] ルーティング設定（`/memos/:id`）
- [ ] モックデータで詳細ページを作成
- [ ] API 呼び出し関数（`GET /memos/:id`）を作成
- [ ] API 連携して実際のデータを表示

### 3. メモ新規作成ページ (`/memos/new`)
- [ ] ルーティング設定（`/memos/new`）
- [ ] モックで新規作成フォームを作成
- [ ] API 呼び出し関数（`POST /memos`）を作成
- [ ] API 連携してメモ作成機能を実装

### 4. メモ編集ページ (`/memos/:id/edit`)
- [ ] ルーティング設定（`/memos/:id/edit`）
- [ ] モックで編集フォームを作成
- [ ] API 呼び出し関数（`PUT /memos/:id`, `DELETE /memos/:id`）を作成
- [ ] API 連携して編集・削除機能を実装

### 5. 仕上げ
- [ ] 最低限のスタイルとレスポンシブ対応

---

## ドキュメント
- [x] `backend-API/docs`にSwaggerでAPIドキュメントを追加する
- [x] `docs/spec-screen-flow.md` に画面遷移図と情報設計概要を追加する

---

## 完了基準
- ローカル環境で Backend と Frontend の主要機能（一覧・詳細・作成・更新・削除）が動作すること
- 必要なドキュメント（スキーマ・API 例）がリポジトリにあること
