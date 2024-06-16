package main

import (
	"context"
	// "net/http"

	"kyewboard/pkg/db"
	"kyewboard/pkg/view"

	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewPlayer() db.Player {
    stats := map[string]int{
        "Vitality": 0,
        "Strength": 0,
        "Inteligence": 0,
        "Sense": 0,
        "Agility": 0,
    }
    dev := db.Skill{Title: "Development", Category: "IT", Level: 1, Experience: 1,}
    sec := db.Skill{Title: "IT Security", Category: "IT", Level: 1, Experience: 1,}
    skate := db.Skill{Title: "Skateboarding", Category: "Sport", Level: 1, Experience: 1,}
    garden := db.Skill{Title: "Gardening", Category: "Biology", Level: 1, Experience: 1,}
    rocketleauge := db.Skill{Title: "Rocketleague", Category: "Esport", Level: 1, Experience: 1,}
    
    skills := map[string]db.Skill{
        "Development": dev,
        "IT Security": sec,
        "Skateboarding": skate,
        "Gardening": garden,
        "Rocketleague": rocketleauge,
    }
    return db.Player{
        Stats : stats,
        Skills: skills,
        Experience : 0,
        Level : 1,
        Id : 1,
        Name : "Kyew",
    }
}


func main() {
	e := echo.New()
    e.Use(middleware.Logger())

    rewards := []string{"+1000 GO Exp", "+1000 Html Exp"}
    objectives := []string{"Setup GO Server", "Setup Templ", "Setup Air", "testing multiline long objecttive omg hi new line wtattat" }
    quest := db.Quest{Id: 1, Message: "Kyewboard setup quest", Status: "Pending",Objectives: objectives, Rewards: rewards, Assignee: "kyew"}
    player := NewPlayer()
	index := view.Index(quest, player)
    body := view.Body(quest, player)
	// quests := make([]Quest)

	// quests := append(quests, quest)
	
	e.Static("/static", "/assets")


	e.GET("/", func(c echo.Context) error {
        return index.Render(context.Background(), c.Response().Writer)
	})


    e.POST("/completed", func(c echo.Context) error {
        return body.Render(context.Background(), c.Response().Writer)
    })


    e.POST("/toggletask", func(c echo.Context) error {
        checked := c.FormValue("taskcheckbox") == "on"
        objective := c.FormValue("tasklabel")

        if checked {
            tasklbl := view.TaskLabelLT(objective)
            return tasklbl.Render(context.Background(), c.Response().Writer)
        } else {
            tasklbl := view.TaskLabel(objective)
            return tasklbl.Render(context.Background(), c.Response().Writer) 
        }
    })


    e.GET("/quests", func(c echo.Context) error {
        return view.QuestPage(quest).Render(context.Background(), c.Response().Writer)
    })


    e.GET("/skills", func(c echo.Context) error {
        return view.Skills(player).Render(context.Background(), c.Response().Writer)
    })

    e.GET("/status", func(c echo.Context) error {
        return view.Status(player).Render(context.Background(), c.Response().Writer)
    })


	e.Logger.Fatal(e.Start(":42069"))
}
