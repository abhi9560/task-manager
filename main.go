package main

import (
    "log"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/abhi9560/task-manager/database"
    "github.com/abhi9560/task-manager/handlers"
)

func main() {
    database.InitDB()

    r := chi.NewRouter()

    r.Get("/tasks", handlers.GetTasks)
    r.Get("/tasks/{id}", handlers.GetTask)
    r.Post("/tasks", handlers.CreateTask)
    r.Put("/tasks/{id}", handlers.UpdateTask)
    r.Delete("/tasks/{id}", handlers.DeleteTask)

    log.Fatal(http.ListenAndServe(":8090", r))
}

