package main

import (
	"context"

	"kyewboard/pkg/db"
	"kyewboard/pkg/view"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	quest := db.Quest{Id: 1, Message: "WRITE A GO SERVER USING TEMPL TAIL WIND AND HTMX", Status: "Pending", Reward: "+ 1000 GO-Exp", Assignee: "kyew"}

	component := view.Index(quest)
	// quests := make([]Quest)

	// quests := append(quests, quest)

	e.Static("/static", "/assets")

	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
