package main

import (
	"database/sql"
	_ "github.com/glebarez/go-sqlite"
	"log"
	"net/http"
)

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatalf("Error setting up database: %s", err)
	}

	// リポジトリを作成する
	repo := &repository{db: db}
	studyHandler := &StudyHandler{repo: repo}

	mux := http.NewServeMux()

	// ハンドラを登録する
	mux.HandleFunc("/first_lessons", studyHandler.GetLastLessonByCourse)

	if err := http.ListenAndServe("127.0.0.1:8080", mux); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func setupDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(`CREATE TABLE courses (course_id INTEGER PRIMARY KEY, name TEXT)`); err != nil {
		return nil, err
	}
	if _, err := db.Exec(`INSERT INTO courses (course_id, name) VALUES (1, "Go"), (2, "Database")`); err != nil {
		return nil, err
	}
	if _, err := db.Exec(`CREATE TABLE lessons (lesson_id INTEGER PRIMARY KEY, course_id INTEGER, created_at TIMESTAMP, name TEXT)`); err != nil {
		return nil, err
	}
	if _, err := db.Exec(`INSERT INTO lessons (lesson_id, course_id, created_at, name)
VALUES
  (1, 1, '2023-01-01', "Go1"),
  (2, 1, '2023-01-02', "Go2"),
  (3, 2, '2023-03-02', "Database1"),
  (4, 2, '2023-05-10', "Database3"),
  (5, 2, '2023-03-10', "Database2")
`); err != nil {
		return nil, err
	}

	return db, err
}
