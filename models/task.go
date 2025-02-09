package models

import (
    "time"
)

type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
    DueDate     time.Time `json:"due_date"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    Cost        float64   `json:"cost"`
    Currency    string    `json:"currency"`
}

type TaskCreate struct {
    Title       string    `json:"title"`
    Description string    `json:"description"`
    DueDate     time.Time `json:"due_date"`
    Cost        float64   `json:"cost"`
    Currency    string    `json:"currency"`
}

type TaskUpdate struct {
    Title       *string    `json:"title,omitempty"`
    Description *string    `json:"description,omitempty"`
    Status      *string    `json:"status,omitempty"`
    DueDate     *time.Time `json:"due_date,omitempty"`
    Cost        *float64   `json:"cost,omitempty"`
    Currency    *string    `json:"currency,omitempty"`
}