package core

// Comment represents a comment on a post
type Comment struct {
	CommentID int
	PostID    int
	UserID    int
	Content   string
}

type Like struct {
	ID     int
	PostID int
	UserID int
}

// Dislike represents a dislike made by a user on a post
type DisLike struct {
	ID     int
	PostID int
	UserID int
}

// Post represents an article published by a user
type Post struct {
	PostID     int
	UserID     int
	CategoryID int
	Title      string
	Content    string
}

// User represents a user
type User struct {
	UserID       int
	Email        string
	Name         string
	PasswordHash string
}

// Category represents a category in which posts can be classified
type Category struct {
	CategoryID  int
	Name        string
	Description string
}
