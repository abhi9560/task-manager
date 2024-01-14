package models

import "time"

type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Priority    int       `json:"priority"`
    DueAt       time.Time `json:"due_at"`
    Day         string    `json:"day"`
}

