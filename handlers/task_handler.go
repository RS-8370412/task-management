package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "fmt"
    "time"
    "task-management/models"
    "task-management/services"
    "github.com/gorilla/mux"
)

type TaskHandler struct {
    taskService     *services.TaskService
    emailService    *services.EmailService
    exchangeService *services.ExchangeService
}

func NewTaskHandler(ts *services.TaskService, es *services.EmailService, xs *services.ExchangeService) *TaskHandler {
    return &TaskHandler{
        taskService:     ts,
        emailService:    es,
        exchangeService: xs,
    }
}

// Home handler for root endpoint
func (h *TaskHandler) Home(w http.ResponseWriter, r *http.Request) {
    message := map[string]string{
        "message": "Welcome to Task Management API",
        "version": "1.0",
        "status": "running",
    }
    json.NewEncoder(w).Encode(message)
}

// CreateTask - Create a new task
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.TaskCreate
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := h.taskService.CreateTask(task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send email notification
    err = h.emailService.SendTaskNotification(
        "rupesh.singh.8370412@gmail.com",
        "New Task Created",
        fmt.Sprintf("New task created: %s", task.Title),
    )
    if err != nil {
        fmt.Printf("Failed to send email: %v\n", err)
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "id": id,
        "message": "Task created successfully",
    })
}

// GetTasks - Get all tasks with optional filtering
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
    status := r.URL.Query().Get("status")
    dueDate := r.URL.Query().Get("due_date")
    
    tasks, err := h.taskService.GetTasks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Filter by status if provided
    if status != "" {
        filteredTasks := []models.Task{}
        for _, task := range tasks {
            if task.Status == status {
                filteredTasks = append(filteredTasks, task)
            }
        }
        tasks = filteredTasks
    }

    // Filter by due date if provided
    if dueDate != "" {
        date, err := time.Parse("2006-01-02", dueDate)
        if err == nil {
            filteredTasks := []models.Task{}
            for _, task := range tasks {
                if task.DueDate.Format("2006-01-02") == date.Format("2006-01-02") {
                    filteredTasks = append(filteredTasks, task)
                }
            }
            tasks = filteredTasks
        }
    }

    // Handle currency conversion if requested
    targetCurrency := r.URL.Query().Get("currency")
    if targetCurrency != "" {
        for i := range tasks {
            convertedCost, err := h.exchangeService.ConvertCurrency(
                tasks[i].Cost,
                tasks[i].Currency,
                targetCurrency,
            )
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            tasks[i].Cost = convertedCost
            tasks[i].Currency = targetCurrency
        }
    }

    json.NewEncoder(w).Encode(tasks)
}

// GetTaskByID - Get a specific task by ID
func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    task, err := h.taskService.GetTaskByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(task)
}

// UpdateTask - Update task details
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    var task models.TaskUpdate
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = h.taskService.UpdateTask(id, task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})
}

// DeleteTask - Delete a task
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    err = h.taskService.DeleteTask(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}

// UpdateTaskStatus - Update task status
func (h *TaskHandler) UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    var update struct {
        Status string `json:"status"`
    }
    if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.taskService.UpdateTaskStatus(id, update.Status); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send email notification
    err = h.emailService.SendTaskNotification(
        "rupesh.singh.8370412@gmail.com",
        "Task Status Updated",
        fmt.Sprintf("Task %d status updated to: %s", id, update.Status),
    )
    if err != nil {
        fmt.Printf("Failed to send email: %v\n", err)
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task status updated successfully"})
}

func (h *TaskHandler) SearchTasks(w http.ResponseWriter, r *http.Request) {
    keyword := r.URL.Query().Get("q")
    if keyword == "" {
        http.Error(w, "Keyword is required", http.StatusBadRequest)
        return
    }

    tasks, err := h.taskService.SearchTasks(keyword)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) ExportTasks(w http.ResponseWriter, r *http.Request) {
    csvData, err := h.taskService.ExportTasks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/csv")
    w.Header().Set("Content-Disposition", "attachment; filename=tasks.csv")
    w.Write(csvData)
}

func (h *TaskHandler) BulkDeleteTasks(w http.ResponseWriter, r *http.Request) {
    // Set response content type
    w.Header().Set("Content-Type", "application/json")

    // Check if request body is empty
    if r.Body == nil {
        http.Error(w, "Request body is required", http.StatusBadRequest)
        return
    }

    var ids []int
    err := json.NewDecoder(r.Body).Decode(&ids)
    if err != nil {
        http.Error(w, "Invalid request body format. Expected array of integers", http.StatusBadRequest)
        return
    }

    if len(ids) == 0 {
        http.Error(w, "No task IDs provided", http.StatusBadRequest)
        return
    }

    err = h.taskService.BulkDeleteTasks(ids)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error deleting tasks: %v", err), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": fmt.Sprintf("Successfully deleted %d tasks", len(ids)),
    })
}