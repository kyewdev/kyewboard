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
        Quests: quests,
    }
}


func main() {
	e := echo.New()
    e.Use(middleware.Logger())

    rewards := []string{"+1000 GO Exp", "+1000 Html Exp"}
    objectives := []db.Objective{
        {Done: true, Text: "Setup GO Server"},
        {Text: "Setup Templ", Done: true},
        {Text: "Setup Air", Done: true},
        {Text: " PUT RL ON M2 ", Done: false}, 
    }
    quest := db.Quest{Id: 1, Message: "Kyewboard setup quest", Status: "Pending",Objectives: objectives, Rewards: rewards, Assignee: "kyew"}
    
    rewards2 := []string{"+1000 DB Exp", "+1000 Docker Exp"}
    objc := []db.Objective{
        {Text: "INSTALL POSTGRE DB ", Done: false}, 
        {Text: "INSTALL DOCKER DESKTOP", Done: false}, 
    }
    quest2 := db.Quest{Id: 2, Message: "PostgreSQL setup quest", Status: "Pending", Objectives: objc, Rewards: rewards2, Assignee: "kyew"}
    
    quests := []db.Quest{quest, quest2}

    player := NewPlayer(quests)

    index := view.Index(quest, player)
    body := view.Body(quest, player)

	
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
        // NEED QEUST UND OBJECTIVE ID

        if checked {
            tasklbl := view.TaskLabelLT(objective)
            return tasklbl.Render(context.Background(), c.Response().Writer)
        } else {
            tasklbl := view.TaskLabel(objective)
            return tasklbl.Render(context.Background(), c.Response().Writer) 
        }
    })


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
