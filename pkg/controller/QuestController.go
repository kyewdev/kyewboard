package controller

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kyewboard/pkg/db"
	"kyewboard/pkg/models"
	"kyewboard/pkg/view"
	"log"
	"net/http"
	"strconv"
)

type QuestController struct {
	Database *gorm.DB
}

func NewQuestController(database *gorm.DB) *QuestController {
	return &QuestController{
		Database: database,
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
	quest, err := db.GetQuestById(qc.Database, questId)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	
	if quest.Status == "Pending" {
		for _, reward := range quest.Rewards {
			reward.Skill.Experience += reward.Amount
			db.SaveEntity(reward.Skill, qc.Database)
		}
		quest.Status = "Done"
			
	} else {
		for _, reward := range quest.Rewards {
			reward.Skill.Experience -= reward.Amount
			db.SaveEntity(reward.Skill, qc.Database)
		}
		quest.Status = "Pending"
	}

	db.SaveEntity(quest, qc.Database)
	return view.Quest(*quest).Render(context.Background(), c.Response().Writer)
}

func (qc *QuestController) ToggleTask(c echo.Context) error {
	checked := c.FormValue("taskcheckbox") == "on"
	objectiveId := c.FormValue("tasklabel")
	objective, err := db.GetObjectiveByID(qc.Database, objectiveId)

	if err != nil {
		log.Printf("Couldnt get Objective with Id: %s, %d", objectiveId, err)
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
    player := db.GetPlayerById(qc.Database, 1)
	return view.EditableReward(player.Skills).Render(context.Background(), c.Response().Writer)
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
	dberr := db.DeleteQuestByID(qc.Database, questIdint)
	if dberr != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Invalid quest_id: %v", dberr))
	}

	return c.NoContent(http.StatusOK)
}

func (qc *QuestController) AddQuest(c echo.Context) error {

	if err := c.Request().ParseForm(); err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse form data")
	}

    quests, qerr := db.GetPendingQuests(qc.Database)
    if qerr != nil {
        return c.NoContent(http.StatusNoContent)
    }

	rewardStrings := c.Request().Form["editableReward"]
    // TODO --> GET SKILLS AS COMBOBOX
    //TODO GET AMOUNT as NR?
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
		Message:    title,
		Status:     "Pending",
		Objectives: objectives,
		Rewards:    rewards,
		Assignee:   "kyew",
	}

	db.SaveEntity(&newQuest, qc.Database)


	return view.QuestPage(quests).Render(context.Background(), c.Response().Writer)
}
