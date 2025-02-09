package services

import (
    "log"
    "bytes"
    "encoding/csv"
    "strconv"
    "fmt"
    "database/sql"
    "errors"
    "task-management/models"
)

type TaskService struct {
    db *sql.DB
}

func NewTaskService(db *sql.DB) *TaskService {
    return &TaskService{db: db}
}

func (s *TaskService) CreateTask(task models.TaskCreate) (int64, error) {
    result, err := s.db.Exec(`
        INSERT INTO tasks (title, description, due_date, cost, currency)
        VALUES (?, ?, ?, ?, ?)`,
        task.Title, task.Description, task.DueDate, task.Cost, task.Currency,
    )
    if err != nil {
        return 0, err
    }
    
    return result.LastInsertId()
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
    rows, err := s.db.Query(`
        SELECT id, title, description, status, due_date, created_at, updated_at, cost, currency
        FROM tasks
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var task models.Task
        err := rows.Scan(
            &task.ID, &task.Title, &task.Description, &task.Status,
            &task.DueDate, &task.CreatedAt, &task.UpdatedAt, &task.Cost, &task.Currency,
        )
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}

func (s *TaskService) GetTaskByID(id int) (*models.Task, error) {
    var task models.Task
    err := s.db.QueryRow(`
        SELECT id, title, description, status, due_date, created_at, updated_at, cost, currency
        FROM tasks
        WHERE id = ?`,
        id,
    ).Scan(
        &task.ID, &task.Title, &task.Description, &task.Status,
        &task.DueDate, &task.CreatedAt, &task.UpdatedAt, &task.Cost, &task.Currency,
    )
    
    if err == sql.ErrNoRows {
        return nil, errors.New("task not found")
    }
    if err != nil {
        return nil, err
    }
    
    return &task, nil
}

func (s *TaskService) UpdateTask(id int, update models.TaskUpdate) error {
    // Get existing task
    existingTask, err := s.GetTaskByID(id)
    if err != nil {
        return err
    }

    // Update fields if provided
    if update.Title != nil {
        existingTask.Title = *update.Title
    }
    if update.Description != nil {
        existingTask.Description = *update.Description
    }
    if update.Status != nil {
        existingTask.Status = *update.Status
    }
    if update.DueDate != nil {
        existingTask.DueDate = *update.DueDate
    }
    if update.Cost != nil {
        existingTask.Cost = *update.Cost
    }
    if update.Currency != nil {
        existingTask.Currency = *update.Currency
    }

    // Update in database
    _, err = s.db.Exec(`
        UPDATE tasks 
        SET title = ?, description = ?, status = ?, due_date = ?, cost = ?, currency = ?
        WHERE id = ?`,
        existingTask.Title, existingTask.Description, existingTask.Status,
        existingTask.DueDate, existingTask.Cost, existingTask.Currency, id,
    )
    return err
}

func (s *TaskService) UpdateTaskStatus(id int, status string) error {
    _, err := s.db.Exec(`
        UPDATE tasks 
        SET status = ?
        WHERE id = ?`,
        status, id,
    )
    return err
}

func (s *TaskService) DeleteTask(id int) error {
    result, err := s.db.Exec(`
        DELETE FROM tasks
        WHERE id = ?`,
        id,
    )
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return errors.New("task not found")
    }

    return nil
}

func (s *TaskService) SearchTasks(keyword string) ([]models.Task, error) {
    rows, err := s.db.Query(`
        SELECT id, title, description, status, due_date, created_at, updated_at, cost, currency
        FROM tasks
        WHERE title LIKE ? OR description LIKE ?`,
        "%"+keyword+"%", "%"+keyword+"%",
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var task models.Task
        err := rows.Scan(
            &task.ID, &task.Title, &task.Description, &task.Status,
            &task.DueDate, &task.CreatedAt, &task.UpdatedAt, &task.Cost, &task.Currency,
        )
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}


func (s *TaskService) ExportTasks() ([]byte, error) {
    tasks, err := s.GetTasks()
    if err != nil {
        log.Printf("Error fetching tasks: %v\n", err) // Log the error
        return nil, err
    }

    log.Printf("Fetched %d tasks for export\n", len(tasks)) // Log the number of tasks

    var csvData bytes.Buffer
    writer := csv.NewWriter(&csvData)

    // Write CSV header
    writer.Write([]string{"ID", "Title", "Description", "Status", "Due Date", "Cost", "Currency"})

    // Write task data
    for _, task := range tasks {
        writer.Write([]string{
            strconv.Itoa(task.ID),
            task.Title,
            task.Description,
            task.Status,
            task.DueDate.Format("2006-01-02 15:04:05"),
            fmt.Sprintf("%.2f", task.Cost),
            task.Currency,
        })
    }

    writer.Flush()
    if err := writer.Error(); err != nil {
        log.Printf("Error writing CSV data: %v\n", err) // Log the error
        return nil, err
    }

    return csvData.Bytes(), nil
}

func (s *TaskService) BulkDeleteTasks(ids []int) error {
    if len(ids) == 0 {
        return fmt.Errorf("no task IDs provided")
    }

    query := "DELETE FROM tasks WHERE id IN ("
    for i, id := range ids {
        if i > 0 {
            query += ","
        }
        query += strconv.Itoa(id)
    }
    query += ")"

    log.Printf("Executing query: %s\n", query)

    _, err := s.db.Exec(query)
    if err != nil {
        log.Printf("Error executing query: %v\n", err)
        return err
    }

    return nil
}