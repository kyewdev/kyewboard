package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kyewboard/pkg/models"
	"log"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=kyewroot dbname=kyewboard port=8181 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, err
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Quest{}, &models.Objective{}, &models.Reward{}, &models.Player{}, &models.Skill{}, &models.PlayerQuest{})
}

func SavePlayer(player models.Player, database *gorm.DB) {

	result := database.Save(&player)

	if result.Error != nil {
		log.Fatalf("failed to save player: %v", result.Error)
	} else {
		log.Printf("PLAYER SAVED; AFFECTED ROWS: %v", result.RowsAffected)
	}

}

func RetrievePlayer(db *gorm.DB, playerID int) (*models.Player, error) {
	var player models.Player
	if err := db.Preload("Skills").Preload("Quests").Preload("Quests.Objectives").Preload("Quests.Rewards").First(&player, playerID).Error; err != nil {
		return nil, err
	}
	return &player, nil
}
