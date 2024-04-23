package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLastLessonByCourse(t *testing.T) {
	// ここにテストを追加する

	// テスト用のデータベースをセットアップする
	db, err := setupDB()
	if err != nil {
		t.Errorf("Error setting up database: %s", err)
	}
	repo := &repository{db: db}
	studyHandler := &StudyHandler{repo: repo}
	// テスト用のサーバーをセットアップする
	ts := httptest.NewServer(http.HandlerFunc(studyHandler.GetLastLessonByCourse))
	defer ts.Close()
	// テストケースを定義する
	testCases := []struct {
		name     string
		expected string
	}{
		{
			name:     "test",
			expected: `[{"LessonID":2,"Course":{"CourseID":1,"Name":"Go"},"Name":"Go2"},{"LessonID":4,"Course":{"CourseID":2,"Name":"Database"},"Name":"Database3"}]`,
		},
	}
	// テストを実行する
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// テスト用のサーバーにリクエストを送信する
			cli := ts.Client()
			resp, err := cli.Get(ts.URL)
			if err != nil {
				t.Errorf("Error getting response: %s", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("expected 200, but got %d", resp.StatusCode)
			}
			// レスポンスボディを取得する
			resBody, err := io.ReadAll(resp.Body)
			// レスポンスボディを検証する
			actual := string(resBody)
			if actual != tc.expected {
				t.Errorf("expected %s, but got %s", tc.expected, actual)
			}
		})
	}
}
