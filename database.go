package main

import (
	"fmt"
	"database/sql"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

type ID_TYPE string
const (
	ROOM ID_TYPE = "room" 
	TASK         = "task"
	USER         = "user"
)

func GetNextID(db *sql.DB, IDType ID_TYPE) (id int) {
	var col string;
	row := db.QueryRow(fmt.Sprintf(`SELECT %s_id FROM %s_names ORDER BY %s_id DESC LIMIT 1`, IDType, IDType, IDType))
	err := row.Scan(&col)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0
		} else {
			panic(err)
		}
	}

	if id, err = strconv.Atoi(col); err != nil {
		panic(err)
	}

	return id + 1
}

func AddNew(db *sql.DB, name string, IDType ID_TYPE) (id int) {
	id = GetNextID(db, IDType)
	_, err := db.Exec(fmt.Sprintf(`INSERT INTO %s_names (%s_id, %s_name) VALUES (%d, '%s')`, IDType, IDType, IDType, id, name))
	if err != nil {
		panic(err)
	}
	return
}

func UpdateTask(db *sql.DB, roomID int, userID int, taskID int, taskType int, progress int) {
	_, err := db.Exec(fmt.Sprintf(`
	INSERT INTO tasks
	(room_id, user_id, task_id, task_type, task_progress)
	VALUES
	(%d, %d, %d, %d, %d)
	ON DUPLICATE KEY UPDATE
	room_id = VALUES(room_id),
	user_id = VALUES(user_id),
	task_id = VALUES(task_id),
	task_type = VALUES(task_type),
	task_progress = VALUES(task_progress)`,
	roomID, userID, taskID, taskType, progress))
	if err != nil {
		panic(err)
	}
	return
}

func AddComment(db *sql.DB, roomID int, userID int, msg string) {
	_, err := db.Exec(fmt.Sprintf(`INSERT INTO comments (room_id, user_id, msg) VALUES (%d, %d, '%s')`, roomID, userID, msg))
	if err != nil {
		panic(err)
	}
}

type Task struct {
	uName string
	tName string
	taskType int
	taskProg int
}

type Comment struct {
	uName string
	msg string
}

func QueryRoom(db *sql.DB, roomID int) ([]Task, []Comment) {


	/// HANDLE TASK QUERY
	rows, err := db.Query(fmt.Sprintf(`SELECT room_id, user_id, task_id, task_type, task_progress FROM tasks WHERE room_id = %d`, roomID))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var tasks []Task;
	var comments []Comment;

	var r, u, t, tp, pr int;
	var uName, tName string;

	for rows.Next() {
		
		rows.Scan(&r, &u, &t, &tp, &pr)

		row := db.QueryRow(fmt.Sprintf(`SELECT user_name FROM user_names WHERE user_id = %d`, u))
		err := row.Scan(&uName)
	
		if err != nil {
			panic(err)
		}

		row = db.QueryRow(fmt.Sprintf(`SELECT task_name FROM task_names WHERE task_id = %d`, t))
		err = row.Scan(&tName)

		if err != nil {
			panic(err)
		}

		tasks = append(tasks, Task{uName, tName, tp, pr})
		// fmt.Printf("%d, %s, %s, %d, %d\n", r, uName, tName, tp, pr);
	}
	
	if err := rows.Err(); err != nil {
		panic(err)
	}


	/// HANDLE COMMENT QUERY

	rows, err = db.Query(fmt.Sprintf(`SELECT room_id, user_id, msg FROM comments WHERE room_id = %d`, roomID))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var msg string;

	for rows.Next() {
		
		rows.Scan(&r, &u, &msg)

		row := db.QueryRow(fmt.Sprintf(`SELECT user_name FROM user_names WHERE user_id = %d`, u))
		err := row.Scan(&uName)
	
		if err != nil {
			panic(err)
		}

		comments = append(comments, Comment{uName, msg})
		// fmt.Printf("%d, %s, %s\n", r, uName, msg);
	}
	
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return tasks, comments;
}

func main() {
	db, err := sql.Open("mysql", "radinyn:password@/appdb")
	if err != nil {
	  panic(err)
	}
}
