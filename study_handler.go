package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	// コースは学習コンテンツの集まりを表す
	Course struct {
		CourseID int
		Name     string
	}

	// Lesson は学習コンテンツを表す
	Lesson struct {
		LessonID int
		Course   *Course
		Name     string
	}

	Repository interface {
		// GetLastLessonByCourse はコースごとの最後に追加された学習コンテンツを取得する
		GetLastLessonByCourse(ctx context.Context) ([]Lesson, error)
	}

	StudyHandler struct {
		repo Repository
	}
)

// GetLastLessonByCourse はコースごとの最後に追加された学習コンテンツを取得する
func (h *StudyHandler) GetLastLessonByCourse(w http.ResponseWriter, r *http.Request) {
	// ここに実装を追加する。
	// リポジトリからデータを取得し、JSONに変換して返す

	// リポジトリからデータを取得
	lessons, err := h.repo.GetLastLessonByCourse(r.Context())
	if err != nil {
		panic(err)
	}
	// JSONに変換してレスポンスを返す
	bytes, err := json.Marshal(lessons)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
	//panic("not implemented")
}
