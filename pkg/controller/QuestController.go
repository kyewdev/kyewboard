package controller

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"kyewboard/pkg/db"
	"kyewboard/pkg/models"
	"kyewboard/pkg/view"
	"log"
	"net/http"
	"strconv"
	"gorm.io/gorm"
)

type QuestController struct {
	Database    *gorm.DB
	PlayerModel *models.Player
}

func NewQuestController(database *gorm.DB, playerModel *models.Player) *QuestController {
	return &QuestController{
		Database:    database,
		PlayerModel: playerModel,
	}
}

func (qc *QuestController) RegisterRoutes(e *echo.Echo) {
	e.POST("/quests/complete", qc.CompleteQuest)
	e.POST("/quests/toggletask", qc.ToggleTask)
	e.GET("/quests/getEditableReward", qc.GetEditableReward)
	e.GET("/quests/getEditableObjective", qc.GetEditableObjective)
	e.DELETE("/quests/delete", qc.DeleteQuest)
	e.POST("/quests/add", qc.AddQuest)
}


func (qc *QuestController) CompleteQuest(c echo.Context) error {
	questId := c.FormValue("questId")
	quest := qc.PlayerModel.GetQuestById(questId)
	if quest == nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return view.QuestPage(qc.PlayerModel.Quests).Render(context.Background(), c.Response().Writer)
}

func (qc *QuestController) ToggleTask(c echo.Context) error {
	checked := c.FormValue("taskcheckbox") == "on"
	objectiveId := c.FormValue("tasklabel")
	objective, err := db.GetObjectiveByID(qc.Database, objectiveId)

	if err != nil {
		log.Fatalf("Couldnt get Objective with Id: %s, %d", objectiveId, err)
	}
	objective.Done = checked
	db.SaveEntity(*objective, qc.Database)

	if checked {
		tasklbl := view.TaskLabelLT(objective.Text)
		return tasklbl.Render(context.Background(), c.Response().Writer)
	} else {
		tasklbl := view.TaskLabel(objective.Text)
		return tasklbl.Render(context.Background(), c.Response().Writer)
	}
}

func (qc *QuestController) GetEditableReward(c echo.Context) error {
	return view.EditableReward().Render(context.Background(), c.Response().Writer)
}

func (qc *QuestController) GetEditableObjective(c echo.Context) error {
	return view.EditableObjective().Render(context.Background(), c.Response().Writer)
}

func (qc *QuestController) DeleteQuest(c echo.Context) error {
	questId := c.FormValue("questId")
	questIdint, err := strconv.Atoi(questId)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Invalid quest_id: %v", err))
	}
	log.Printf("WOULD DELETE QUEST: %d", questIdint)
	qc.PlayerModel.RemoveQuestByID(questIdint)
    for _, q := range qc.PlayerModel.Quests {
        log.Printf(q.Message)
    }
    db.DeleteQuestByID(qc.Database, questIdint) 
//	db.SaveEntity(*qc.PlayerModel, qc.Database)
    
	return view.QuestPage(qc.PlayerModel.Quests).Render(context.Background(), c.Response().Writer)
}



func (qc *QuestController) AddQuest(c echo.Context) error {
	if err := c.Request().ParseForm(); err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse form data")
	}

	rewardStrings := c.Request().Form["editableReward"]
	if len(rewardStrings) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	rewards := []models.Reward{}
	for _, r := range rewardStrings {
		if r != "" {
			reward := models.Reward{Text: r}
			rewards = append(rewards, reward)
		}
	}

	objectiveStrings := c.Request().Form["editableObjective"]
	if len(objectiveStrings) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	objectives := []models.Objective{}
	for _, obj := range objectiveStrings {
		if obj != "" {
			objective := models.Objective{Text: obj, Done: false}
			objectives = append(objectives, objective)
		}
	}

	title := c.FormValue("editableTitle")

	newQuest := models.Quest{
		ID:         len(qc.PlayerModel.Quests) + 1,
		Message:    title,
		Status:     "Pending",
		Objectives: objectives,
		Rewards:    rewards,
		Assignee:   "kyew",
	}
	qc.PlayerModel.Quests = append(qc.PlayerModel.Quests, newQuest)
	db.SaveEntity(*qc.PlayerModel, qc.Database)

	return view.QuestPage(qc.PlayerModel.Quests).Render(context.Background(), c.Response().Writer)
}
