package models

import (
	"strconv"
)

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

type Player struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	Level      int
	Experience int
	Skills     []Skill        `gorm:"foreignKey:PlayerID;constraint:OnDelete:CASCADE"`
	Stats      map[string]int `gorm:"-"`
	Quests     []Quest        `gorm:"many2many:player_quests;constraint:OnDelete:CASCADE"`
}

func (p *Player) GetQuestById(questId string) *Quest {

	for _, quest := range p.Quests {
		if strconv.Itoa(quest.ID) == questId {
			return &quest
		}
	}
	return nil
}

func (p *Player) RemoveQuestByID(questID int) {
	for i, quest := range p.Quests {
		if quest.ID == questID {
			p.Quests = append(p.Quests[:i], p.Quests[i+1:]...)
			break
		}
	}
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
