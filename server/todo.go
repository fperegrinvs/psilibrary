package main

import "time"

// Todo objeto exemplo CRUD
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos lista de objetos Todo
type Todos []Todo