package main

import (
	"context"
	// "net/http"

	"kyewboard/pkg/db"
	"kyewboard/pkg/view"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewPlayer(quests []db.Quest) db.Player {
	stats := map[string]int{
		"Vitality":    0,
		"Strength":    0,
		"Inteligence": 0,
		"Sense":       0,
		"Agility":     0,
	}
	dev := db.Skill{Title: "Development", Category: "IT", Level: 1, Experience: 1}
	sec := db.Skill{Title: "IT Security", Category: "IT", Level: 1, Experience: 1}
	skate := db.Skill{Title: "Skateboarding", Category: "Sport", Level: 1, Experience: 1}
	garden := db.Skill{Title: "Gardening", Category: "Biology", Level: 1, Experience: 1}
	rocketleauge := db.Skill{Title: "Rocketleague", Category: "Esport", Level: 1, Experience: 1}

	skills := map[string]db.Skill{
		"Development":   dev,
		"IT Security":   sec,
		"Skateboarding": skate,
		"Gardening":     garden,
		"Rocketleague":  rocketleauge,
	}
	return db.Player{
		Stats:      stats,
		Skills:     skills,
		Experience: 0,
		Level:      1,
		ID:         1,
		Name:       "Kyew",
		Quests:     quests,
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	rewards1 := []string{"+1000 GO Exp", "+1000 Html Exp"}
	objc1 := []db.Objective{
		{Done: true, Text: "Setup GO Server"},
		{Text: "Setup Templ", Done: true},
		{Text: "Setup Air", Done: true},
		{Text: " PUT RL ON M2 ", Done: false},
	}
	quest := db.Quest{ID: 1, Message: "Kyewboard setup quest", Status: "Pending", Objectives: objc1, Rewards: rewards1, Assignee: "kyew"}

	rewards2 := []string{"+1000 DB Exp", "+1000 Docker Exp"}
	objc2 := []db.Objective{
		{Text: "INSTALL POSTGRE DB ", Done: false},
		{Text: "INSTALL DOCKER DESKTOP", Done: false},
	}
	quest2 := db.Quest{ID: 2, Message: "PostgreSQL setup quest", Status: "Pending", Objectives: objc2, Rewards: rewards2, Assignee: "kyew"}

	rewards3 := []string{"+1000 Game Dev. Exp", "+1000 C++ Exp"}
	objc3 := []db.Objective{
		{Done: false, Text: "Setup Project"},
	}

	quest3 := db.Quest{ID: 3, Message: "Kyewgame Setup Quest", Status: "Pending", Objectives: objc3, Rewards: rewards3, Assignee: "kyew"}
	quests := []db.Quest{quest, quest2, quest3}

	player := NewPlayer(quests)

	index := view.Index(quest, player)

	/////////////BASE //////////////////
	e.Static("/static", "/assets")

	e.GET("/", func(c echo.Context) error {
		return index.Render(context.Background(), c.Response().Writer)
	})

	////////////////QUEST /////////////////////////
	e.POST("/completed", func(c echo.Context) error {
		return view.QuestPage(quests).Render(context.Background(), c.Response().Writer)
	})

	e.POST("/toggletask", func(c echo.Context) error {
		checked := c.FormValue("taskcheckbox") == "on"
		objective := c.FormValue("tasklabel")
		// NEED QEUST UND OBJECTIVE ID

		if checked {
			tasklbl := view.TaskLabelLT(objective)
			return tasklbl.Render(context.Background(), c.Response().Writer)
		} else {
			tasklbl := view.TaskLabel(objective)
			return tasklbl.Render(context.Background(), c.Response().Writer)
		}
	})

	e.POST("/addquest", func(c echo.Context) error {
		// GET OBJEVTIVES AND REWARDS -> RETURN NEW QUEST PAGE WITH n + 1 quests
		reward := c.FormValue("editableReward")
		rewards := []string{reward}
		objective := c.FormValue("editableObjective")
		objectives := []db.Objective{
			{Done: false, Text: objective},
		}
		title := c.FormValue("editableTitle")
		newQuest := db.Quest{ID: len(quests) + 1, Message: title, Status: "Pending", Objectives: objectives, Rewards: rewards, Assignee: "kyew"}
		quests = append(quests, newQuest)
		return view.QuestPage(quests).Render(context.Background(), c.Response().Writer)
	})
	//////////// PAGES /////////////////////////
	e.GET("/quests", func(c echo.Context) error {
		return view.QuestPage(quests).Render(context.Background(), c.Response().Writer)
	})

	e.GET("/skills", func(c echo.Context) error {
		return view.Skills(player).Render(context.Background(), c.Response().Writer)
	})

	e.GET("/status", func(c echo.Context) error {
		return view.Status(player).Render(context.Background(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
