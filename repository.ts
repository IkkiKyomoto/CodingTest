import { Repository, Lesson } from './study_handler'

export interface Database {
  all(sql: string, ...args: any[]): Promise<any[]>
}

export class RepositoryImpl implements Repository {
  constructor(private db: Database) { }

  // GetLastLessonByCourse はコースごとの最後に追加された学習コンテンツを取得する
  async GetLastLessonByCourse(): Promise<Lesson[]> {
    // ここに実装を追加する。
    // データベースからデータを取得し、Lessonのスライスを返す
    // レッスンのみではなく、コースも取得する必要がある
    throw new Error('Not implemented')
  }
}
