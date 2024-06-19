package core

import (
	"database/sql"
	"log"
	//"strconv"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func InitDatabase(dbname string) *sql.DB {
	db, err := sql.Open("sqlite3", "./database/"+dbname+".db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS User (
			UserID INTEGER PRIMARY KEY,
			Email VARCHAR(255) UNIQUE,
			Name VARCHAR(100),
			PasswordHash VARCHAR(255)
		);
		CREATE TABLE IF NOT EXISTS Category (
			CategoryID INTEGER PRIMARY KEY,
			Name VARCHAR(100),
			Description TEXT
		);
		CREATE TABLE IF NOT EXISTS Posts (
			PostID INTEGER PRIMARY KEY,
			UserID INTEGER NOT NULL,
			CategoryID INTEGER NOT NULL,
			Title VARCHAR(255),
			Content TEXT,
			FOREIGN KEY (UserID) REFERENCES User(UserID),
			FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID)
		);
		CREATE TABLE IF NOT EXISTS Like (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			PostID INTEGER NOT NULL,
			UserID INTEGER NOT NULL,
			FOREIGN KEY (PostID) REFERENCES Posts(PostID),
			FOREIGN KEY (UserID) REFERENCES User(UserID)
		);
		CREATE TABLE IF NOT EXISTS Dislike (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			PostID INTEGER NOT NULL,
			UserID INTEGER NOT NULL,
			FOREIGN KEY (PostID) REFERENCES Posts(PostID),
			FOREIGN KEY (UserID) REFERENCES User(UserID)
		);
		CREATE TABLE IF NOT EXISTS Comments (
			CommentID INTEGER PRIMARY KEY AUTOINCREMENT,
			PostID INTEGER NOT NULL,
			UserID INTEGER NOT NULL,
			Content TEXT,
			FOREIGN KEY (PostID) REFERENCES Posts(PostID),
			FOREIGN KEY (UserID) REFERENCES User(UserID)
		);
		PRAGMA foreign_keys = ON;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	return db
}

func InserttoUsers(db *sql.DB, name string, email string, password string) (int64, error) {
	query1 := `INSERT INTO User (Name, Email, PasswordHash) VALUES (?, ?, ?)`
	stmt, err := db.Prepare(query1)
	if err != nil {
		log.Printf("%q: %s\n", err, query1)
		return 0, nil
	}
	defer stmt.Close()

	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Printf("Password Hashing Error: %v", err)
		return 0, err
	}

	result, err := stmt.Exec(name, email, hashedPassword)
	if err != nil {
		log.Printf("%q: %s\n", err, query1)
		return 0, nil
	}
	return result.LastInsertId()
}

func SelectAllFromUsers(db *sql.DB) []User {
	query := `SELECT UserID, Email, Name, PasswordHash FROM User`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return nil
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserID, &user.Email, &user.Name, &user.PasswordHash)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		users = append(users, user)
	}
	return users
}

func SelectUserById(db *sql.DB, id int) User {
	query := `SELECT UserID, Email, Name, PasswordHash FROM User WHERE UserID = ?`
	row := db.QueryRow(query, id)
	var user User
	err := row.Scan(&user.UserID, &user.Email, &user.Name, &user.PasswordHash)
	if err != nil {
		log.Fatalf("Scan: %v", err)
	}
	return user
}

func SelectUserNameWithPattern(db *sql.DB, pattern string) []User {
	query := `SELECT UserID, Email, Name, PasswordHash FROM User WHERE Name LIKE ?`
	rows, err := db.Query(query, "%"+pattern+"%")
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return nil
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserID, &user.Email, &user.Name, &user.PasswordHash)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		users = append(users, user)
	}
	return users
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Post Management

// SelectPostById selects a post by its ID from the Posts table.
func SelectPostById(db *sql.DB, id int) Post {
	query := `SELECT PostID, UserID, CategoryID, Title, Content FROM Posts WHERE PostID = ?`
	row := db.QueryRow(query, id)
	var post Post
	err := row.Scan(&post.PostID, &post.UserID, &post.CategoryID, &post.Title, &post.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No post found with ID %d", id)
		} else {
			log.Fatalf("Scan: %v", err)
		}
	}
	return post
}

// SelectPostByUser selects all posts by a specified user from the Posts table.
func SelectPostByUser(db *sql.DB, userID int) []Post {
	query := `SELECT PostID, UserID, CategoryID, Title, Content FROM Posts WHERE UserID = ?`
	rows, err := db.Query(query, userID)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return nil
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.PostID, &post.UserID, &post.CategoryID, &post.Title, &post.Content)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		posts = append(posts, post)
	}
	return posts
}

// InsertIntoContent inserts a new post into the Posts table.
func InsertIntoContent(db *sql.DB, title string, content string, category_id, user_id int) (int64, error) {
	query := `INSERT INTO Posts (UserID, CategoryID, Title, Content) VALUES (?, ?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user_id, category_id, title, content)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, err
	}
	return result.LastInsertId()
}

// DeletePostFromId deletes a post by its ID from the Posts table.
func DeletePostFromId(db *sql.DB, id int) (int64, error) {
	query := `DELETE FROM Posts WHERE PostID = ?`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected: %v", err)
		return 0, err
	}
	return rowsAffected, nil
}

// Likes, Dislikes, and Other Reactions Management

// SelectLikeById2 selects a like by PostID and UserID from the Like table.
func SelectLikeById2(db *sql.DB, postID, userID int) Like {
	query := `SELECT ID, PostID, UserID FROM Like WHERE PostID = ? AND UserID = ?`
	row := db.QueryRow(query, postID, userID)
	var like Like
	err := row.Scan(&like.ID, &like.PostID, &like.UserID)
	if err != nil {
		log.Fatalf("Scan: %v", err)
	}
	return like
}

func InsertIntoLike(db *sql.DB, postID, userID int) (int64, error) {
	query := `INSERT INTO Like (PostID, UserID) VALUES (?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, nil
	}
	defer stmt.Close()

	result, err := stmt.Exec(postID, userID)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, nil
	}
	return result.LastInsertId()
}

// InsertIntoDislike inserts a new dislike into the Dislike table.
func InsertIntoDislike(db *sql.DB, postID, userID int) (int64, error) {
	query := `INSERT INTO Dislike (PostID, UserID) VALUES (?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(postID, userID)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, err
	}
	return result.LastInsertId()
}

// DeleteLikeFromId deletes a like (or other reaction) by its ID from the specified table.
func DeleteLikeFromId(db *sql.DB, table string, id int) (int64, error) {
	query := `DELETE FROM ` + table + ` WHERE ID = ?`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected: %v", err)
		return 0, err
	}
	return rowsAffected, nil
}

// InsertComment inserts a new comment into the Comments table.
func InsertComment(db *sql.DB, postID, userID int, content string) (int64, error) {
	query := `INSERT INTO Comments (PostID, UserID, Content) VALUES (?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(postID, userID, content)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, err
	}
	return result.LastInsertId()
}

// SelectCommentsByPostID selects all comments for a specified post.
func SelectCommentsByPostID(db *sql.DB, postID int) []Comment {
	query := `SELECT CommentID, PostID, UserID, Content FROM Comments WHERE PostID = ?`
	rows, err := db.Query(query, postID)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return nil
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.Content)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		comments = append(comments, comment)
	}
	return comments
}

// DeleteCommentByID deletes a comment by its ID from the Comments table.
func DeleteCommentByID(db *sql.DB, id int) (int64, error) {
	query := `DELETE FROM Comments WHERE CommentID = ?`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected: %v", err)
		return 0, err
	}
	return rowsAffected, nil
}
