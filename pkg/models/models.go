package models

type Player struct {
    ID         int `gorm:"primaryKey;autoIncrement"`
    Name       string
    Level      int
    Experience int
    Skills     []Skill        `gorm:"foreignKey:PlayerID;constraint:OnDelete:CASCADE"`
    Stats      map[string]int `gorm:"-"`
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

type Quest struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Message    string
	Status     string
	Objectives []Objective `gorm:"foreignKey:QuestID;constraint:OnDelete:CASCADE;"`
	Rewards    []Reward    `gorm:"foreignKey:QuestID;constraint:OnDelete:CASCADE;"`
	Assignee   string
	Questtype  string
	Category   string
}

type Reward struct {
	ID      int `gorm:"primaryKey;autoIncrement"`
	Text    string
	Amount  int
	SkillID int `gorm:"index"`
	QuestID int `gorm:"index"`
	Skill   Skill
}

type Objective struct {
	ID      int `gorm:"primaryKey;autoIncrement"`
	Done    bool
	Text    string
	QuestID int `gorm:"index"`
}


