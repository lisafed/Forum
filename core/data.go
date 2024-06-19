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

func InsertIntoUsers(db *sql.DB, name string, email string, password string) (int64, error) {
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

func SelectLikeById(db *sql.DB, postID int) []Like {
	query := `SELECT ID, PostID, UserID FROM Like WHERE PostID = ?`
	rows, err := db.Query(query, postID)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
		return nil
	}
	defer rows.Close()

	likes := []Like{}
	for rows.Next() {
		var like Like
		err = rows.Scan(&like.ID, &like.PostID, &like.UserID)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		likes = append(likes, like)
	}
	return likes
}

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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
