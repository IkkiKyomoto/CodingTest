# コーディングテスト

本リポジトリは、TRUNK 株式会社の採用選考のためのコーディングテストです。
目安の時間は50分ですが、それ以上の時間をかけて頂いても構いません。
言語はGoかTypeScriptのいずれかを選択してください。
以下、Goを選択した場合の説明です。

## 課題

以下の3ステップに従って、学習サービスのサーバーの一部APIを実装してください。
ステップごとにGitのコミットを行い、コミットログを残してください。

### ステップ1

`study_handler.go` の `GetLastLessonByCourse` 関数を実装してください。
リポジトリからデータを取得し、JSON形式に変換して返すようにしてください。

### ステップ2

`repository.go` の `GetLastLessonByCourse` 関数を実装してください。
本ステップが完了すると、以下のようにAPIを実行することができます。

```
% go run ./main.go
```

```
% curl http://127.0.0.1:8080/first_lessons
[{"LessonID":2,"Course":{"CourseID":1,"Name":"Go"},"Name":"Go2"},{"LessonID":4,"Course":{"CourseID":2,"Name":"Database"},"Name":"Database3"}]%
```

### ステップ3

`study_handler_test.go` の `TestGetLastLessonByCourse` 関数を実装してください。

`go test` を実行して、テストが通ることを確認してください。

### TypeScriptを選択した場合

同じファイル名の.tsファイルがありますので、TypeScriptを選択した場合はそちらを参照してください。
ビルド、実行、テストの方法は、以下になります。

```bash
# ビルド
% npm run build

# 実行
% npm run start

# テスト
% npm run test
```

### 注意事項

- リポジトリ内のその他のファイルやコードは変更しないでください。
- 何を考えて実装したか、コード内にコメントで記述してください。

### 提出方法

解いたリポジトリをご自身のGitHubにアップロードするか、`zip` で圧縮して、`dev@trunk.school` 宛にURLかZIPデータを送信してください。

## 採点基準

- コードの品質 (Go,TypeScriptのベストプラクティスに従っているかは見ません)
- SQLの効率性
- テスト観点

## お問い合わせ先

何かご不明点があれば、`dev@trunk.school` までお問い合わせください。
