package main

type IDType string

type User struct {
	ID   IDType `json:"id"`
	Name string `json:"name"`
}

type SmallTask struct {
	ID   IDType `json:"id"`
	Name string `json:"name"`
}

type ComplexTask struct {
	ID   IDType `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID      IDType `json:"id"`
	UserID  IDType `json:"user_id"`
	Conetnt string `json:"content"`
}

type Room struct {
	ID           IDType   `json:"id"`
	Users        []IDType `json:"users"`
	SmallTasks   []IDType `json:"smallTasks"`
	ComplexTasks []IDType `json:"complexTasks"`
}
