package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/n0rdb4hnh0f/GoBBS-API/models"
)

// ValidatePost は投稿内容をチェックする（再利用可能にする）
func validatePost(post *models.Post) string {
	if post.Content == "" {
		return "コンテンツは必須です"
	}
	if len(post.Content) > 1000 { // 現実的な数値に変更
		return "コンテンツが長すぎます"
	}
	if post.ThreadID == "" {
		return "スレッドIDが指定されていません"
	}
	if post.Author == "" {
		post.Author = "名無しさん"
	}
	if len(post.Author) > 50 {
		return "名前が長すぎます"
	}
	return ""
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		ErrorJSON(w, http.StatusBadRequest, "リクエスト形式が正しくありません")
		return
	}

	if errMsg := validatePost(&post); errMsg != "" {
		ErrorJSON(w, http.StatusBadRequest, errMsg)
		return
	}

	// DB保存
	if err := models.DB.WithContext(r.Context()).Create(&post).Error; err != nil {
		ErrorJSON(w, http.StatusInternalServerError, "保存に失敗しました")
		return
	}

	ResponseJSON(w, http.StatusCreated, post)
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	if err := models.DB.Order("created_at desc").Limit(100).Find(&posts).Error; err != nil {
		ErrorJSON(w, http.StatusInternalServerError, "取得に失敗しました")
		return
	}
	ResponseJSON(w, http.StatusOK, posts)
}
