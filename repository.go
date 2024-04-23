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
	// SQL文について：　まず、lessonとcourseをcourse_idで結合(A)させ、さらに、サブクエリで各コースごとに最後に追加されたレッスンを取得(B)し、AとBを自然結合させる。この操作によって、各コースごとに最後に追加されたレッスンを取得する
	rows, err := r.db.QueryContext(ctx, "SELECT l.lesson_id, c.course_id, l.name, c.name FROM (lessons as l INNER JOIN courses as c ON l.course_id = c.course_id) NATURAL INNER JOIN (SELECT lesson_id, max(created_at) AS last_created FROM lessons GROUP BY course_id)")
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
	//　取得したデータを返す
	return lessons, nil
}
