package models
import(
)


type Quest struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Message    string
	Status     string
	Objectives []Objective `gorm:"foreignKey:QuestID"`
	Rewards    []Reward    `gorm:"foreignKey:QuestID"`
	Assignee   string
	Questtype  string
	Category   string
}

type Reward struct {
	ID      int `gorm:"primaryKey;autoIncrement"`
	Type    string
	Text    string
	Amount  int
	QuestID int `gorm:"index"`
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
