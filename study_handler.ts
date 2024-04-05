export type Course = {
  course_id: number;
  name: string;
}

export type Lesson = {
  lesson_id: number;
  course_id: number;
  name: string;
}

export interface Repository {
  GetLastLessonByCourse(): Promise<Lesson[]>;
}

export class StudyHandler {
  constructor(private repo: Repository) { }

  // GetLastLessonByCourse はコースごとの最後に追加された学習コンテンツを取得する
  async GetLastLessonByCourse(): Promise<Lesson[]> {
    // ここに実装を追加する。
    // リポジトリからデータを取得し、JSONに変換して返す
    throw new Error('Not implemented')
  }
}
