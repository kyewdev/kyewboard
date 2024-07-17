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
		log.Fatalf("failed to connect to the database: %v", retrErr)
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

func PlayerWithQuests() models.Player {
	rewards1 := []models.Reward{
		{Text: "+1000 GO Exp"},
		{Text: "+1000 Html Exp"},
	}
	objc1 := []models.Objective{
		{Done: true, Text: "Setup GO Server"},
		{Text: "Setup Templ", Done: true},
		{Text: "Setup Air", Done: true},
		{Text: " PUT RL ON M2 ", Done: false},
	}

	quest := models.Quest{ID: 1, Message: "Kyewboard setup quest", Status: "Pending", Objectives: objc1, Rewards: rewards1, Assignee: "kyew"}

	rewards2 := []models.Reward{{Text: "+1000 DB Exp"}, {Text: "+1000 Docker Exp"}}
	objc2 := []models.Objective{
		{Text: "INSTALL POSTGRE DB ", Done: false},
		{Text: "INSTALL DOCKER DESKTOP", Done: false},
	}
	quest2 := models.Quest{ID: 2, Message: "PostgreSQL setup quest", Status: "Pending", Objectives: objc2, Rewards: rewards2, Assignee: "kyew"}

	rewards3 := []models.Reward{
		{Text: "+1000 Game Dev. Exp"},
		{Text: "+1000 C++ Exp"}}
	objc3 := []models.Objective{
		{Done: false, Text: "Setup Project"},
	}

	quest3 := models.Quest{ID: 3, Message: "Kyewgame Setup Quest", Status: "Pending", Objectives: objc3, Rewards: rewards3, Assignee: "kyew"}
	quests := []models.Quest{quest, quest2, quest3}

	return NewPlayer(quests)

}

func NewPlayer(quests []models.Quest) models.Player {
	stats := map[string]int{
		"Vitality":    0,
		"Strength":    0,
		"Inteligence": 0,
		"Sense":       0,
		"Agility":     0,
	}
	dev := models.Skill{Title: "Development", Category: "IT", Level: 1, Experience: 1}
	sec := models.Skill{Title: "IT Security", Category: "IT", Level: 1, Experience: 1}
	skate := models.Skill{Title: "Skateboarding", Category: "Sport", Level: 1, Experience: 1}
	garden := models.Skill{Title: "Gardening", Category: "Biology", Level: 1, Experience: 1}
	rocketleauge := models.Skill{Title: "Rocketleague", Category: "Esport", Level: 1, Experience: 1}

	skills := []models.Skill{
		dev,
		sec,
		skate,
		garden,
		rocketleauge,
	}
	return models.Player{
		Stats:      stats,
		Skills:     skills,
		Experience: 0,
		Level:      1,
		ID:         1,
		Name:       "Kyew",
		Quests:     quests,
	}
}
