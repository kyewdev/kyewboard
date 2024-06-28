package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	return db.AutoMigrate(&Quest{}, &Objective{}, &Player{}, &Skill{}, &PlayerQuest{})
}

type Quest struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Message    string
	Status     string
	Objectives []Objective `gorm:"foreignKey:QuestID"`
	Rewards    []string    `gorm:"type:text[]"`
	Assignee   string
	Questtype  string
	Category   string
}

type Objective struct {
	ID      int `gorm:"primaryKey;autoIncrement"`
	Done    bool
	Text    string
	QuestID int `gorm:"index"`
}

type Player struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	Level      int
	Experience int
	Skills     []Skill        `gorm:"foreignKey:PlayerID"`
	Stats      map[string]int `gorm:"-"`
	Quests     []Quest        `gorm:"many2many:player_quests;"`
}

type Skill struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	Category    string
	Level       int
	Experience  int
	Title       string
	Description string
	PlayerID    int `gorm:"index"`
}

type PlayerQuest struct {
	PlayerID int `gorm:"primaryKey"`
	QuestID  int `gorm:"primaryKey"`
}
