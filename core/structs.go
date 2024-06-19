package core

// Comment représente un commentaire sur un post
type Comment struct {
	CommentID int    // Identifiant unique du commentaire
	PostID    int    // Identifiant du post sur lequel le commentaire est fait
	UserID    int    // Identifiant de l'utilisateur qui a fait le commentaire
	Content   string // Contenu du commentaire
}

// like représente un like fait par un utilisateur sur un post
type Like struct {
	ID     int // Identifiant unique du like/dislike
	PostID int // Identifiant du post concerné
	UserID int // Identifiant de l'utilisateur qui a effectué l'action
}

// Dislike représente un dislike  fait par un utilisateur sur un post
type DisLike struct {
	ID     int // Identifiant unique du like/dislike
	PostID int // Identifiant du post concerné
	UserID int // Identifiant de l'utilisateur qui a effectué l'action
}

// Post représente un article publié par un utilisateur
type Post struct {
	PostID     int    // Identifiant unique du post
	UserID     int    // Identifiant de l'utilisateur qui a créé le post
	CategoryID int    // Identifiant de la catégorie du post
	Title      string // Titre du post
	Content    string // Contenu du post
}

// User représente un utilisateur
type User struct {
	UserID       int    // Identifiant unique de l'utilisateur
	Email        string // Email de l'utilisateur, doit être unique
	Name         string // Nom de l'utilisateur
	PasswordHash string // Hash du mot de passe de l'utilisateur
}

// Category représente une catégorie dans laquelle les posts peuvent être classés
type Category struct {
	CategoryID  int    // Identifiant unique de la catégorie
	Name        string // Nom de la catégorie
	Description string // Description de la catégorie
}
