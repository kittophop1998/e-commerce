package infrastructures

import (
	"log"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	entities := RegisterEntities()
	for _, entity := range entities {
		if err := db.AutoMigrate(entity); err != nil {
			log.Fatalf("Error migrating %v: %v", entity, err)
		}
	}

	log.Println("Migration completed")
}
