package main

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

// GetLastLessonByCourse はコースごとの最後に追加された学習コンテンツを取得する
func (r *repository) GetLastLessonByCourse(ctx context.Context) ([]Lesson, error) {
	// ここに実装を追加する。
	// データベースからデータを取得し、Lessonのスライスを返す
	// レッスンのみではなく、コースも取得する必要がある
	panic("not implemented")
}
