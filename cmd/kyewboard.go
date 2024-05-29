package main

import (
	"context"
	"net/http"

	"kyewboard/pkg/db"
	"kyewboard/pkg/view"

	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewPlayer() db.Player {
    skills := map[string]int{
        "Vitality": 0,
        "Strength": 0,
        "Inteligence": 0,
        "Sense": 0,
        "Agility": 0,
    }

    return db.Player{
        Skills : skills,
        Experience : 0,
        Level : 1,
        Id : 1,
        Name : "Kyew",
    }
}


func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	quest := db.Quest{Id: 1, Message: "WRITE A GO SERVER USING TEMPL TAIL WIND AND HTMX", Status: "Pending", Reward: "+ 1000 GO-Exp", Assignee: "kyew"}
    player := NewPlayer()
	component := view.Index(quest, player)
	// quests := make([]Quest)

	// quests := append(quests, quest)
	
	e.Static("/static", "/assets")

	e.GET("/", func(c echo.Context) error {
        return component.Render(context.Background(), c.Response().Writer)
	})

    e.POST("/accepted", func(c echo.Context) error {
        quest.Status = "Accepted"
        completebtn := view.CompleteDiv()
        return completebtn.Render(context.Background(),c.Response().Writer)

    })

    e.POST("/declined", func(c echo.Context) error {
        quest.Status = "Declined"
        return c.String(http.StatusOK, quest.Status)
    })

	e.Logger.Fatal(e.Start(":42069"))
}
