package main

import "time"

// Post представляет собой структуру данных для поста в блоге
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

// Посты блога будут храниться в срезе
var posts []Post
var nextID int = 1
