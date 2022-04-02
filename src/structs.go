package main

type IDType string

type Comment struct {
	UserID  IDType `json:"user_id"`
	Conetnt string `json:"content"`
}

type Room struct {
	ID           IDType                 `json:"id"`
	Users        []string               `json:"users"`
	SmallTasks   map[string]([]bool)    `json:"smallTasks"`
	ComplexTasks map[string]([]float32) `json:"complexTasks"`
	Comments     []Comment              `json:"comments"`
}
