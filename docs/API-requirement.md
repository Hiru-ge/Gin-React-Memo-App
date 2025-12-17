# API仕様書

## メモ一覧取得
- URL: /memos
- Method: GET
- Response: 200 OK
``` JSON
[
    {
        "id":1,
        "title":"買い物リスト",
        "content":"牛乳、卵、パン",
        "created_at": "2023-10-27T10:00:00Z" 
    },
    {
        "id":2,
        "title":"買い物リスト2",
        "content":"牛乳、卵、パン、鶏むね肉",
        "created_at": "2023-10-28T10:00:00Z"  
    }
]
```

## メモ詳細取得
- URL: /memos/:id
- Method: GET
- Response: 200 OK
``` JSON
[
    {
        "id":1,
        "title":"買い物リスト",
        "content":"牛乳、卵、パン",
        "created_at": "2023-10-27T10:00:00Z" 
    }
]
```

## 新規メモ作成
- URL: /memos
- Method: POST
- Request:
``` JSON
[
    {
        "title":"買い物リスト",
        "content":"牛乳、卵、パン",
    }
]
```

- Response: 201 Created
``` JSON
[
    {
        "id":1,
        "title":"新しいメモ",
        "content":"メモの内容です",
        "created_at": "2023-10-27T10:00:00Z" 
    }
]
```

## メモ編集
- URL: /memos
- Method: PUT
- Request:
``` JSON
[
    {
        "title":"更新後タイトル",
        "content":"更新後の内容です",
    }
]
```

- Response: 200 OK (更新後の内容を返す)
``` JSON
[
    {
        "id":1,
        "title":"更新後タイトル",
        "content":"更新後の内容です",
        "created_at": "2023-10-27T10:00:00Z" 
    }
]
```

## メモ削除
- URL: /memos/:id
- Method: DELETE
- Response: 204 NoContent (ボディ無し)
