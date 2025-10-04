package main

import (
	"context"
	"kyewboard/pkg/controller"
	"kyewboard/pkg/db"
	"kyewboard/pkg/models"
	"kyewboard/pkg/view"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var playermodel *models.Player

	e.Use(middleware.Logger())

	database, connErr := db.Connect()
	if connErr != nil {
		log.Fatalf("failed to connect to the database: %v", connErr)
	}

	if migErr := db.Migrate(database); migErr != nil {
		log.Fatalf("failed to migrate database: %v", migErr)
	}

	playermodel = db.GetPlayerById(database, 1)

	if playermodel == nil {
		playermodel =  NewPlayer()
		db.SaveEntity(&playermodel, database)
		skills := NewSkills()
		quests := NewQuests()
		playermodel.Skills = skills
		db.SaveEntity(&playermodel, database)
        db.SaveEntity(&quests, database)
		log.Printf("setup skills and quests")
	}
	qc := controller.NewQuestController(database)
	qc.RegisterRoutes(e)
	index := view.Index(*playermodel)

	/////////////BASE //////////////////
	e.Static("/static", "assets")

	e.GET("/", func(c echo.Context) error {
		return index.Render(context.Background(), c.Response().Writer)
	})

	//////////// PAGES /////////////////////////
	e.GET("/quests", func(c echo.Context) error {
        var quests []models.Quest
        quests, err := db.GetPendingQuests(database)
		if err != nil {
            log.Printf("Couldnt retrieve pending quests QUESTS: %v", err)
		}
		return view.QuestPage(quests).Render(context.Background(), c.Response().Writer)
	})

	e.GET("/skills", func(c echo.Context) error {
		playermodel := db.GetPlayerById(database, 1)

		if playermodel == nil {
			log.Printf("couldt load player for skills")
			return c.NoContent(http.StatusInternalServerError)
		}

		return view.Skills(*playermodel).Render(context.Background(), c.Response().Writer)
	})

	e.GET("/status", func(c echo.Context) error {
		return view.Status(*playermodel).Render(context.Background(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":42069"))
}

func NewPlayer() *models.Player {
	stats := map[string]int{
		"Vitality":    0,
		"Strength":    0,
		"Inteligence": 0,
		"Sense":       0,
		"Agility":     0,
	}
	return &models.Player{
		Stats:      stats,
		Experience: 0,
		Level:      1,
		ID:         1,
		Name:       "Kyew",
	}
}

func NewSkills() []models.Skill {
	return []models.Skill{
		{ID: 1, Title: "Development", Category: "IT", Level: 1, Experience: 1, },
		{ID: 2, Title: "IT Security", Category: "IT", Level: 1, Experience: 1, },
		{ID: 3, Title: "Skateboarding", Category: "Sport", Level: 1, Experience: 1, },
		{ID: 4, Title: "Gardening", Category: "Biology", Level: 1, Experience: 1, },
		{ID: 5, Title: "Rocketleague", Category: "Esport", Level: 1, Experience: 1, },
	}
}

func NewQuests() []models.Quest {
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
				{Text: "+400 Dev Exp", SkillID: 1, Amount: 400},
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
				{Text: "+400 Dev Exp", SkillID: 1, Amount: 400},
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
				{Text: "+500 Game Dev. Exp", SkillID: 1, Amount: 500},
			},
		},
	}
}
