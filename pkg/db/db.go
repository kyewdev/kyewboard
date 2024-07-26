package db

import (
	"fmt"
	"kyewboard/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=kyewroot dbname=kyewboard port=8181 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

	if err != nil {
		return nil, err
	}
	return db, err
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Quest{}, &models.Objective{}, &models.Reward{}, &models.Player{}, &models.Skill{}, &models.PlayerQuest{})
}

func SaveEntity(entity interface{}, database *gorm.DB) (error) {
	result := database.Session(&gorm.Session{FullSaveAssociations: true}).Save(entity)

	if result.Error != nil {
		log.Printf("failed to save entity: %v", result.Error)
	} else {
		log.Printf("ENTITY SAVED; AFFECTED ROWS: %v", result.RowsAffected)
	}
    return result.Error
}
func GetPlayerById(db *gorm.DB, playerID int) (*models.Player, error) {
	var player models.Player
	if err := db.Preload("Skills").Preload("Quests").Preload("Quests.Objectives").Preload("Quests.Rewards").First(&player, playerID).Error; err != nil {
		return nil, err
	}
	return &player, nil
}

func GetQuestById(db *gorm.DB, questID string) (*models.Quest, error) {
    var quest models.Quest
	if err := db.Preload("Objectives").Preload("Rewards").Preload("Rewards.Skill").First(&quest, questID).Error; err != nil {
		return nil, err
	}
    return &quest, nil
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


func GetSkillByID(database *gorm.DB, skillID string) (*models.Skill, error) {
	var skill models.Skill
	if err := database.First(&skill, skillID).Error; err != nil {
		return nil, err
	}
	return &skill, nil
}

func DeletePlayerByID(db *gorm.DB, playerID int) error {
	if err := db.Delete(&models.Player{}, playerID).Error; err != nil {
		return fmt.Errorf("failed to delete player: %w", err)
	}
	return nil
}


func DeleteQuestByID(db *gorm.DB, questID int) error {
	if err := db.Delete(&models.Quest{}, questID).Error; err != nil {
		return fmt.Errorf("failed to delete player: %w", err)
	}
	return nil
}
