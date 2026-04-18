# GoBBS-API
Go製の掲示板API

GoBBS-APIの仕様をMarkdownでまとめました。フロントエンド実装の参考にしてください。

---

# API ドキュメント

## ベースURL
`http://localhost:8080`

---

## 1. スレッド管理

### スレッド一覧を取得
* **URL:** `/threads`
* **Method:** `GET`
* **Response (JSON):**
```json
[
  {
    "ID": 1,
    "CreatedAt": "2026-04-18T15:00:00Z",
    "title": "最初のスレッド"
  }
]
```

### スレッド詳細を取得
* **URL:** `/threads/{id}`
* **Method:** `GET`
* **Response (JSON):**
```json
{
  "ID": 1,
  "title": "最初のスレッド",
  "posts": [
    {
      "ID": 1,
      "thread_id": 1,
      "content": "こんにちは",
      "author": "名無しさん"
    }
  ]
}
```

### スレッドを作成
* **URL:** `/threads`
* **Method:** `POST`
* **Request (JSON):**
```json
{
  "title": "新しいスレッド"
}
```

---

## 2. 投稿管理

### 投稿一覧を取得
* **URL:** `/posts`
* **Method:** `GET`
* **Response (JSON):**
```json
[
  {
    "ID": 1,
    "thread_id": 1,
    "content": "こんにちは",
    "author": "名無しさん"
  }
]
```

### 投稿を作成
* **URL:** `/posts`
* **Method:** `POST`
* **Request (JSON):**
```json
{
  "thread_id": 1,
  "content": "投稿内容です",
  "author": "ユーザー名"
}
```

---

## エラーレスポンス
バリデーションエラーやシステムエラー時は、以下の形式で返されます。

```json
{
  "status": 400,
  "message": "エラー内容のテキスト"
}
```
