package core

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// UserService handles operations related to users
type UserService struct {
	DB *sql.DB
}

// RegisterUser registers a new user
func (s *UserService) RegisterUser(email, username, password string) error {
	// Check if the email already exists
	exists, err := s.EmailExists(email)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("email already in use")
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Create a new User instance
	newUser := User{
		Email:        email,
		Name:         username,
		PasswordHash: string(hashedPassword),
	}

	// Insert the new user into the User table
	query := `INSERT INTO User (Email, Name, PasswordHash) VALUES (?, ?, ?)`
	_, err = s.DB.Exec(query, newUser.Email, newUser.Name, newUser.PasswordHash)
	if err != nil {
		return fmt.Errorf("failed to insert new user: %w", err)
	}

	return nil
}

// EmailExists checks if an email is already used
func (s *UserService) EmailExists(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM User WHERE Email = ?)`
	err := s.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check email existence: %w", err)
	}
	return exists, nil
}

// LoginUser verifies user credentials
func (s *UserService) LoginUser(emailOrUsername, password string) error {
	var hashedPassword string
	query := `SELECT PasswordHash FROM User WHERE Email = ? OR Name = ?`
	err := s.DB.QueryRow(query, emailOrUsername, emailOrUsername).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return fmt.Errorf("invalid credentials")
	}
	if err != nil {
		return fmt.Errorf("failed to query password hash: %w", err)
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid credentials")
	}
	return nil
}

// ChangeUserPassword allows a user to change their password
func (s *UserService) ChangeUserPassword(userID int, newPassword string) error {
	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Update the user's password in the database
	query := `UPDATE User SET PasswordHash = ? WHERE UserID = ?`
	_, err = s.DB.Exec(query, hashedPassword, userID)
	if err != nil {
		return fmt.Errorf("failed to change user password: %w", err)
	}
	return nil
}

// UpdateUserProfile updates the user's profile information
func (s *UserService) UpdateUserProfile(userID int, email, name string) error {
	query := `UPDATE User SET Email = ?, Name = ? WHERE UserID = ?`
	_, err := s.DB.Exec(query, email, name, userID)
	if err != nil {
		return fmt.Errorf("failed to update user profile: %w", err)
	}
	return nil
}

// DeleteUser deletes a user account
func (s *UserService) DeleteUser(userID int) error {
	query := `DELETE FROM User WHERE UserID = ?`
	_, err := s.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// LogoutUser removes the user session
func (s *UserService) LogoutUser(sessionID string) error {
	query := `DELETE FROM Sessions WHERE SessionID = ?`
	result, err := s.DB.Exec(query, sessionID)
	if err != nil {
		return fmt.Errorf("failed to logout: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no session found with the specified sessionID")
	}
	return nil
}
