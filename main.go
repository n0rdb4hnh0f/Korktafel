package main

import (
	"fmt"
	"net/http"

	"github.com/n0rdb4hnh0f/GoBBS-API/handlers"
	"github.com/n0rdb4hnh0f/GoBBS-API/models"
)

func main() {
	// 1. データベースの初期化
	models.InitDB()

	// 2. ルーティングの設定
	// GETとPOSTを同じパスで処理したい場合は、handler側でメソッド判定するか、
	// ここでクロージャを使って分岐させます。
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetPostsHandler(w, r)
		case http.MethodPost:
			handlers.CreatePostHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running at http://localhost:8080")

	// 3. サーバーの起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
