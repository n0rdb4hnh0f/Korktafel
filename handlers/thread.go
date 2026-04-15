package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/n0rdb4hnh0f/GoBBS-API/models"
	"gorm.io/gorm"
)

func validateThread(thread *models.Thread) string {
	if thread.Title == "" {
		return "スレッドタイトルは必須です"
	}
	if len(thread.Title) > 100 {
		return "タイトルが長すぎます（100文字以内）"
	}
	return ""
}

func CreateThreadHandler(w http.ResponseWriter, r *http.Request) {
	var thread models.Thread
	if err := json.NewDecoder(r.Body).Decode(&thread); err != nil {
		ErrorJSON(w, http.StatusBadRequest, "リクエスト形式が正しくありません")
		return
	}

	if errMsg := validateThread(&thread); errMsg != "" {
		ErrorJSON(w, http.StatusBadRequest, errMsg)
		return
	}

	if err := models.DB.WithContext(r.Context()).Create(&thread).Error; err != nil {
		ErrorJSON(w, http.StatusInternalServerError, "スレッドの作成に失敗しました")
		return
	}

	ResponseJSON(w, http.StatusCreated, thread)
}

func GetThreadsHandler(w http.ResponseWriter, r *http.Request) {
	var threads []models.Thread
	// 最新の作成順に並べる。CreatedAtはBaseに含まれている前提
	if err := models.DB.WithContext(r.Context()).Order("created_at desc").Find(&threads).Error; err != nil {
		ErrorJSON(w, http.StatusInternalServerError, "スレッドの取得に失敗しました")
		return
	}
	ResponseJSON(w, http.StatusOK, threads)
}

func GetThreadDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var thread models.Thread

	err := models.DB.WithContext(r.Context()).
		Preload("Posts", "deleted_at IS NULL").
		Preload("Posts", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at ASC")
		}).
		First(&thread, "id = ?", id).Error

	if err != nil {
		ErrorJSON(w, http.StatusNotFound, "スレッドが見つかりません")
		return
	}

	ResponseJSON(w, http.StatusOK, thread)
}
