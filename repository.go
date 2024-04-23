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

	//　DBからデータを取得　取得の際はレッスンとコースを結合して取得する
	lessons := []Lesson{}
	rows, err := r.db.QueryContext(ctx, "SELECT lessons.lesson_id, courses.course_id, lessons.name, courses.name FROM lessons INNER JOIN courses ON lessons.course_id = courses.course_id")
	if err != nil {
		return nil, err
	}

	//　取得したデータをスライスに格納する
	for rows.Next() {
		//コースについては、一度Course構造体に格納してからLesson構造体に格納する
		var lesson Lesson
		var course Course
		err := rows.Scan(&lesson.LessonID, &course.CourseID, &lesson.Name, &course.Name)
		lesson.Course = &course
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}
