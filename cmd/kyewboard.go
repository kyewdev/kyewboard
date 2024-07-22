package main

import (
	"context"
	"kyewboard/pkg/controller"
	"kyewboard/pkg/db"
	"kyewboard/pkg/models"
	"kyewboard/pkg/view"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	database, connErr := db.Connect()
	if connErr != nil {
		log.Fatalf("failed to connect to the database: %v", connErr)
	}

	if migErr := db.Migrate(database); migErr != nil {
		log.Fatalf("failed to migrate database: %v", migErr)
	}
    

	playermodel, retrErr := db.GetPlayerById(database, 1)
    
	if retrErr != nil {
		log.Printf("failed to connect to the database: %v", retrErr)
        playermodel := NewPlayer()    
        db.SaveEntity(&playermodel, database)
        skills := NewSkillsForPlayer(1)
        quests := NewQuestsForPlayer(1)
        playermodel.Skills = skills
        playermodel.Quests = quests
        db.SaveEntity(&playermodel, database)
        log.Printf("setup skills and quests")
	}
	qc := controller.NewQuestController(database, playermodel)
	qc.RegisterRoutes(e)
	index := view.Index(*playermodel)

	/////////////BASE //////////////////
	e.Static("/static", "/assets")

	e.GET("/", func(c echo.Context) error {
		return index.Render(context.Background(), c.Response().Writer)
	})

	//////////// PAGES /////////////////////////
	e.GET("/quests", func(c echo.Context) error {
		return view.QuestPage(playermodel.Quests).Render(context.Background(), c.Response().Writer)
	})

	e.GET("/skills", func(c echo.Context) error {
		return view.Skills(*playermodel).Render(context.Background(), c.Response().Writer)
	})

	e.GET("/status", func(c echo.Context) error {
		return view.Status(*playermodel).Render(context.Background(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":42069"))
}


func NewPlayer() models.Player {
	stats := map[string]int{
		"Vitality":    0,
		"Strength":    0,
		"Inteligence": 0,
		"Sense":       0,
		"Agility":     0,
	}
	return models.Player{
		Stats:      stats,
		Experience: 0,
		Level:      1,
		ID:         1,
		Name:       "Kyew",
	}
}


func NewSkillsForPlayer(playerID int) []models.Skill {
	return []models.Skill{
		{Title: "Development", Category: "IT", Level: 1, Experience: 1, PlayerID: playerID},
		{Title: "IT Security", Category: "IT", Level: 1, Experience: 1, PlayerID: playerID},
		{Title: "Skateboarding", Category: "Sport", Level: 1, Experience: 1, PlayerID: playerID},
		{Title: "Gardening", Category: "Biology", Level: 1, Experience: 1, PlayerID: playerID},
		{Title: "Rocketleague", Category: "Esport", Level: 1, Experience: 1, PlayerID: playerID},
	}
}

func NewQuestsForPlayer(playerID int) []models.Quest {
	return []models.Quest{
		{
			Message:  "Kyewboard setup quest",
			Status:   "Pending",
			Assignee: "kyew",
			Objectives: []models.Objective{
				{Done: true, Text: "Setup GO Server"},
				{Done: true, Text: "Setup Templ"},
				{Done: true, Text: "Setup Air"},
				{Done: false, Text: "PUT RL ON M2"},
			},
			Rewards: []models.Reward{
				{Text: "+1000 GO Exp"},
				{Text: "+1000 Html Exp"},
			},
		},
		{
			Message:  "PostgreSQL setup quest",
			Status:   "Pending",
			Assignee: "kyew",
			Objectives: []models.Objective{
				{Done: false, Text: "INSTALL POSTGRE DB"},
				{Done: false, Text: "INSTALL DOCKER DESKTOP"},
			},
			Rewards: []models.Reward{
				{Text: "+1000 DB Exp"},
				{Text: "+1000 Docker Exp"},
			},
		},
		{
			Message:  "Kyewgame Setup Quest",
			Status:   "Pending",
			Assignee: "kyew",
			Objectives: []models.Objective{
				{Done: false, Text: "Setup Project"},
			},
			Rewards: []models.Reward{
				{Text: "+1000 Game Dev. Exp"},
				{Text: "+1000 C++ Exp"},
			},
		},
	}
}
