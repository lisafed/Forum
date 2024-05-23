package structs

// LikeDislike représente un like ou dislike fait par un utilisateur sur un post
type LikeDislike struct {
	ID     int    // Identifiant unique du like/dislike
	PostID int    // Identifiant du post concerné
	UserID int    // Identifiant de l'utilisateur qui a effectué l'action
	Type   string // Type d'action ('like' ou 'dislike')
}
