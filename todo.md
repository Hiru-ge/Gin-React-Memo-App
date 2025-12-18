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
- [ ] `database/sql` と `go-sql-driver/mysql` で DB 接続を実装（DSN は環境変数で管理）
- [ ] モデルと CRUD を実装（`GET /memos`, `GET /memos/:id`, `POST /memos`, `PUT /memos/:id`, `DELETE /memos/:id`）
- [ ] Postman / curl で API を手動確認する

---

## Frontend (React + TypeScript)
- [ ] Vite でプロジェクト作成（`react-ts` テンプレート推奨）
- [ ] React Router を導入してルーティングを実装
  - 推奨ルート: `/` (一覧), `/memos/new` (新規作成), `/memos/:id` (詳細), `/memos/:id/edit` (編集)
- [ ] ページ/コンポーネント: メモ一覧、メモ詳細、投稿/編集フォームを作る
- [ ] `src/api` に API 呼び出し関数を作成し UI と連携する
- [ ] 最低限のスタイルとレスポンシブ対応

---

## ドキュメント
- [ ] `docs/API-requirement.md` を実装に合わせて更新
- [ ] 使い方の簡単な curl 例を追加する

---

## 完了基準
- ローカル環境で Backend と Frontend の主要機能（一覧・詳細・作成・更新・削除）が動作すること
- 必要なドキュメント（スキーマ・API 例）がリポジトリにあること
