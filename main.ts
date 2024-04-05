import fastify from 'fastify'
import sqlite3 from 'sqlite3'
import util from 'util'
import { RepositoryImpl, Database } from './repository'
import { StudyHandler } from './study_handler'

async function main() {
  const db = await setupDB()

  const server = fastify()

  const repo = new RepositoryImpl(db)
  const handler = new StudyHandler(repo)

  server.get('/first_lessons', async () => {
    const lessons = handler.GetLastLessonByCourse()

    return lessons
  })

  server.listen({ port: 8080 }, (err, address) => {
    if (err) {
      console.error(err)
      process.exit(1)
    }
    console.log(`Server listening at ${address}`)
  })
}

main().catch((err) => {
  console.error(err)
  process.exit(1)
})

async function setupDB(): Promise<Database> {
  const db = new sqlite3.Database(':memory:');

  const exec = util.promisify(db.exec).bind(db)

  await exec('CREATE TABLE courses (course_id INTEGER PRIMARY KEY, name TEXT)')
  await exec('INSERT INTO courses (course_id, name) VALUES (1, "Go"), (2, "Database")')
  await exec('CREATE TABLE lessons (lesson_id INTEGER PRIMARY KEY, course_id INTEGER, created_at TIMESTAMP, name TEXT)')
  await exec(`INSERT INTO lessons (lesson_id, course_id, created_at, name)
VALUES
  (1, 1, '2023-01-01', "Go1"),
  (2, 1, '2023-01-02', "Go2"),
  (3, 2, '2023-03-02', "Database1"),
  (4, 2, '2023-05-10', "Database3"),
  (5, 2, '2023-03-10', "Database2")
`)

  return new DatabaseImpl(db)
}

class DatabaseImpl implements Database {
  constructor(private db: sqlite3.Database) { }

  async all(sql: string, ...args: any[]): Promise<any[]> {
    return new Promise((resolve, reject) => {
      this.db.all(sql, args, (err, rows) => {
        if (err) {
          reject(err)
          return
        }
        resolve(rows)
      })
    })
  }
}

