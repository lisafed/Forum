package main

import (
	"Forum/data"
	"log"
)

func main() {
	dbPath := "path/to/your/forum.db"
	db := data.initDB(dbPath)
	defer db.Close()

	// Ici, vous pouvez ajouter du code pour manipuler la base de données
	log.Println("Base de données initialisée et prête à être utilisée")
}
