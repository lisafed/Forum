package core

/*
import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)


import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// UserService gère les opérations liées aux utilisateurs
type UserService struct {
	DB *sql.DB
}

// RegisterUser permet d'inscrire un nouvel utilisateur
func (s *UserService) RegisterUser(email, username, password string) error {
	exists, err := s.EmailExists(email)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := User{
		Email:            email,
		Name:             username,
		PasswordHash:     string(hashedPassword),
		RegistrationDate: time.Now(),
	}

	query := `INSERT INTO Users (Email, Name, PasswordHash, RegistrationDate) VALUES (?, ?, ?, ?)`
	_, err = s.DB.Exec(query, newUser.Email, newUser.Name, newUser.PasswordHash, newUser.RegistrationDate)
	if err != nil {
		return fmt.Errorf("failed to insert new user: %w", err)
	}

	return nil
}

// EmailExists vérifie si un email est déjà utilisé
func (s *UserService) EmailExists(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM Users WHERE Email = ?)`
	err := s.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check email existence: %w", err)
	}
	return exists, nil
}

// LoginUser vérifie les identifiants de l'utilisateur
func (s *UserService) LoginUser(emailOrUsername, password string) error {
	var hashedPassword string
	query := `SELECT PasswordHash FROM Users WHERE Email = ? OR Name = ?`
	err := s.DB.QueryRow(query, emailOrUsername, emailOrUsername).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return fmt.Errorf("invalid credentials")
	}
	if err != nil {
		return fmt.Errorf("failed to query password hash: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid credentials")
	}
	return nil
}

// ChangeUserPassword permet de changer le mot de passe d'un utilisateur
func (s *UserService) ChangeUserPassword(userID int, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	query := `UPDATE Users SET PasswordHash = ? WHERE UserID = ?`
	_, err = s.DB.Exec(query, hashedPassword, userID)
	if err != nil {
		return fmt.Errorf("failed to change user password: %w", err)
	}
	return nil
}

// UpdateUserProfile permet de mettre à jour les informations de profil d'un utilisateur
func (s *UserService) UpdateUserProfile(userID int, email, name, profilePicture string) error {
	query := `UPDATE Users SET Email = ?, Name = ?, ProfilePicture = ? WHERE UserID = ?`
	_, err := s.DB.Exec(query, email, name, profilePicture, userID)
	if err != nil {
		return fmt.Errorf("failed to update user profile: %w", err)
	}
	return nil
}

// DeleteUser permet de supprimer un compte utilisateur
func (s *UserService) DeleteUser(userID int) error {
	query := `DELETE FROM Users WHERE UserID = ?`
	_, err := s.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// PromoteToModerator permet de promouvoir un utilisateur au rang de modérateur
func (s *UserService) PromoteToModerator(userID int) error {
	query := `UPDATE Users SET Role = 'moderator' WHERE UserID = ?`
	_, err := s.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to promote user to moderator: %w", err)
	}
	return nil
}

// DemoteToUser permet de rétrograder un modérateur au rang d'utilisateur normal
func (s *UserService) DemoteToUser(userID int) error {
	query := `UPDATE Users SET Role = 'user' WHERE UserID = ?`
	_, err := s.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to demote moderator to user: %w", err)
	}
	return nil
}

// LogoutUser supprime la session utilisateur
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
*/
