package structs

import "time"

// Comment représente un commentaire sur un post
type Comment struct {
	CommentID int       // Identifiant unique du commentaire
	PostID    int       // Identifiant du post sur lequel le commentaire est fait
	UserID    int       // Identifiant de l'utilisateur qui a fait le commentaire
	Content   string    // Contenu du commentaire
	DateTime  time.Time // Date et heure du commentaire
}
