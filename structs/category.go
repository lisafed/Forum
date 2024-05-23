package structs

// Category représente une catégorie dans laquelle les posts peuvent être classés
type Category struct {
	CategoryID  int    // Identifiant unique de la catégorie
	Name        string // Nom de la catégorie
	Description string // Description de la catégorie
}
