package database

import (
    "database/sql"
    "log"
//    "time"
    _ "github.com/lib/pq"
    "github.com/abhi9560/task-manager/models"
)

var DB *sql.DB

// databse connection
func InitDB() {
    var err error
    DB, err = sql.Open("postgres", "postgres://mydb:Cavisson123@localhost/mydb?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
}

// GetTasks prints all task in database
func GetTasks() ([]models.Task, error) {
	rows, err := DB.Query("SELECT id, title, description, priority, due_at, day FROM tasks")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.DueAt, &task.Day)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// GetTaskByID retrieves a specific task by ID from the database.
func GetTaskByID(id int) (*models.Task, error) {
	var task models.Task
	err := DB.QueryRow("SELECT id, title, description, priority, due_at, day FROM tasks WHERE id = $1", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.DueAt, &task.Day)
	if err == sql.ErrNoRows {
		return nil, nil // Task not found
	} else if err != nil {
		log.Println(err)
		return nil, err
	}

	return &task, nil
}

// CreateTask creates a new task and stores it in the database.
func CreateTask(task *models.Task) error {
	_, err := DB.Exec("INSERT INTO tasks (title, description, priority, due_at, day) VALUES ($1, $2, $3, $4, $5)",
		task.Title, task.Description, task.Priority, task.DueAt, task.Day)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// UpdateTask updates an existing task in the database.
func UpdateTask(task *models.Task) error {
	_, err := DB.Exec("UPDATE tasks SET title = $1, description = $2, priority = $3, due_at = $4, day = $5 WHERE id = $6",
		task.Title, task.Description, task.Priority, task.DueAt, task.Day, task.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// DeleteTask deletes a specific task by ID from the database.
func DeleteTask(id int) error {
	_, err := DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

