package structs

import "time"

// Post représente un article publié par un utilisateur
type Post struct {
	PostID     int       // Identifiant unique du post
	UserID     int       // Identifiant de l'utilisateur qui a créé le post
	CategoryID int       // Identifiant de la catégorie du post
	Title      string    // Titre du post
	Content    string    // Contenu du post
	DateTime   time.Time // Date et heure de la publication
	Like       LikeDislike
}
