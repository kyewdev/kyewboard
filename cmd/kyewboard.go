package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"kyewboard/pkg/db"
	"kyewboard/pkg/models"
	"kyewboard/pkg/view"
	"log"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	database, connErr := db.Connect()
	if connErr != nil {
		log.Fatalf("failed to connect to the database: %v", connErr)
	}

	// if migErr := db.Migrate(database); migErr != nil {
	//	log.Fatalf("failed to migrate database: %v", migErr)
	//}

	playermodel, retrErr := db.GetPlayerById(database, 1)

	if retrErr != nil {
		log.Fatalf("failed to connect to the database: %v", retrErr)
	}

	index := view.Index(*playermodel)

	/////////////BASE //////////////////
	e.Static("/static", "/assets")

	e.GET("/", func(c echo.Context) error {
		return index.Render(context.Background(), c.Response().Writer)
	})

	////////////////QUEST /////////////////////////
	e.POST("/quests/complete", func(c echo.Context) error {
		questId := c.FormValue("questId")
		quest := playermodel.GetQuestById(questId)
		if quest == nil {
			return c.NoContent(http.StatusBadRequest)
		}
		return view.QuestPage(playermodel.Quests).Render(context.Background(), c.Response().Writer)
	})

	e.POST("/quests/toggletask", func(c echo.Context) error {
		checked := c.FormValue("taskcheckbox") == "on"
		objectiveId := c.FormValue("tasklabel")
		objective, err := db.GetObjectiveByID(database, objectiveId)

		if err != nil {
			log.Fatalf("Couldnt get Objective with Id: %s, %d", objectiveId, retrErr)
		}
		objective.Done = checked
		db.SaveEntity(*objective, database)
		// NEED QEUST UND OBJECTIVE ID

		if checked {
			tasklbl := view.TaskLabelLT(objective.Text)
			return tasklbl.Render(context.Background(), c.Response().Writer)
		} else {
			tasklbl := view.TaskLabel(objective.Text)
			return tasklbl.Render(context.Background(), c.Response().Writer)
		}
	})

	e.GET("/quests/getEditableReward", func(c echo.Context) error {
		return view.EditableReward().Render(context.Background(), c.Response().Writer)

	})

	e.GET("/quests/getEditableObjective", func(c echo.Context) error {
		return view.EditableObjective().Render(context.Background(), c.Response().Writer)

	})

	e.DELETE("/quests/delete", func(c echo.Context) error {
		questId := c.FormValue("questId")
		questIdint, err := strconv.Atoi(questId)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Invalid quest_id: %v", err))
		}
		log.Printf("WOUD DELETE QUEST: %d", questIdint)
		playermodel.RemoveQuestByID(questIdint)
        db.SaveEntity(*playermodel, database)

		return view.QuestPage(playermodel.Quests).Render(context.Background(), c.Response().Writer)
	})

	e.POST("/quests/add", func(c echo.Context) error {
		// GET OBJEVTIVES AND REWARDS -> RETURN NEW QUEST PAGE WITH n + 1 quests
		// Parse the form values
		if err := c.Request().ParseForm(); err != nil {
			return c.String(http.StatusBadRequest, "Failed to parse form data")
		}

		rewardStrings := c.Request().Form["editableReward"]
		if len(rewardStrings) == 0 {
			return c.NoContent(http.StatusNoContent)
		}


		rewards := []models.Reward{}
		for _, r := range rewardStrings {
            if r != "" { // Skip empty objectives
                reward := models.Reward{
                    Text: r,
                }
                rewards = append(rewards, reward)
            }
		}
        
        objectiveStrings := c.Request().Form["editableObjective"]


		if len(objectiveStrings) == 0 {
			return c.NoContent(http.StatusNoContent)
		}

        objectives := []models.Objective{}
		for _, obj := range objectiveStrings {
            if obj != "" { // Skip empty objectives
                objective := models.Objective{
                    Text: obj,
                    Done: false,
                }
                objectives = append(objectives, objective)
            }
		}

		title := c.FormValue("editableTitle")

		newQuest := models.Quest{ID: len(playermodel.Quests) + 1, Message: title, Status: "Pending", Objectives: objectives, Rewards: rewards, Assignee: "kyew"}
		playermodel.Quests = append(playermodel.Quests, newQuest)
		db.SaveEntity(*playermodel, database)


		return view.QuestPage(playermodel.Quests).Render(context.Background(), c.Response().Writer)
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
