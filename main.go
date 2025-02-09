package main

import (
    "log"
    "net/http"
    "strings"
    "task-management/config"
    "task-management/database"
    "task-management/handlers"
    "task-management/services"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Error loading config:", err)
    }

    // Initialize database
    db, err := database.InitDB(cfg.GetDBConnString())
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }
    defer db.Close()

    // Initialize services
    taskService := services.NewTaskService(db)
    emailService := services.NewEmailService(cfg.SendGridAPIKey)
    exchangeService := services.NewExchangeService(cfg.ExchangeRateAPIKey)

    // Initialize handler
    taskHandler := handlers.NewTaskHandler(taskService, emailService, exchangeService)

    // Set up router
    r := mux.NewRouter()

    // API routes
    api := r.PathPrefix("/api").Subrouter()

    // IMPORTANT: Route order matters! More specific routes must come first
    // Bulk operations and special endpoints
    taskRoutes := api.PathPrefix("/tasks").Subrouter()
    taskRoutes.HandleFunc("/bulk", taskHandler.BulkDeleteTasks).Methods("DELETE")
    taskRoutes.HandleFunc("/export", taskHandler.ExportTasks).Methods("GET")
    taskRoutes.HandleFunc("/search", taskHandler.SearchTasks).Methods("GET")

    // Regular CRUD operations
    taskRoutes.HandleFunc("", taskHandler.CreateTask).Methods("POST")
    taskRoutes.HandleFunc("", taskHandler.GetTasks).Methods("GET")
    taskRoutes.HandleFunc("/{id}/status", taskHandler.UpdateTaskStatus).Methods("PUT")
    taskRoutes.HandleFunc("/{id}", taskHandler.GetTaskByID).Methods("GET")
    taskRoutes.HandleFunc("/{id}", taskHandler.UpdateTask).Methods("PUT")
    taskRoutes.HandleFunc("/{id}", taskHandler.DeleteTask).Methods("DELETE")

    // Static files - should be last
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

    // Log all registered routes
    log.Println("Registered Routes:")
    err = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
        path, err := route.GetPathTemplate()
        if err != nil {
            return err
        }

        methods, err := route.GetMethods()
        if err != nil {
            methods = []string{"*"}
        }

        log.Printf("Route: %-6s %s", strings.Join(methods, ","), path)
        return nil
    })

    if err != nil {
        log.Fatal("Error walking routes:", err)
    }

    // Add CORS middleware
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type", "application/json"},
    })

    handler := c.Handler(r)
    
    log.Println("Server starting on :3090")
    log.Fatal(http.ListenAndServe(":3090", handler))
}
