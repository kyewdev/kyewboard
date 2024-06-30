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

func SaveEntity(entity interface{}, database *gorm.DB) {
	result := database.Save(entity)

	if result.Error != nil {
		log.Fatalf("failed to save entity: %v", result.Error)
	} else {
		log.Printf("ENTITY SAVED; AFFECTED ROWS: %v", result.RowsAffected)
	}
}
func GetPlayerById(db *gorm.DB, playerID int) (*models.Player, error) {
	var player models.Player
	if err := db.Preload("Skills").Preload("Quests").Preload("Quests.Objectives").Preload("Quests.Rewards").First(&player, playerID).Error; err != nil {
		return nil, err
	}
	return &player, nil
}

func GetQuestById(db *gorm.DB, questID int) (*models.Quest, error) {
    var quest models.Quest
	if err := db.Preload("Objectives").Preload("Rewards").First(&quest, questID).Error; err != nil {
		return nil, err
	}
    return &quest,nil
}
func GetObjectiveByID(database *gorm.DB, objectiveID string) (*models.Objective, error) {
	var objective models.Objective
	if err := database.First(&objective, objectiveID).Error; err != nil {
		return nil, err
	}
	return &objective, nil
}

func GetRewardByID(database *gorm.DB, rewardID string) (*models.Reward, error) {
	var reward models.Reward
	if err := database.First(&reward, rewardID).Error; err != nil {
		return nil, err
	}
	return &reward, nil
}
