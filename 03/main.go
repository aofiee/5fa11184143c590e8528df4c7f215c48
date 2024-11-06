package main

import (
	"count-beef/handler"
	"count-beef/repository"
	"count-beef/service"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	uri = "https://baconipsum.com"
	// /api/?type=meat-and-filler&paras=99&format=text
)

func main() {
	app := fiber.New()
	app.Use(fiberLogger.New(fiberLogger.Config{
		Format:     "[${time}] ${method} ${path}",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Bangkok",
	}))

	rp := repository.NewRepo(uri)
	sv := service.NewService(rp)
	hdl := handler.NewHandler(sv)

	bf := app.Group("/beef")
	{
		bf.Get("/summary", hdl.GetSummary)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			err := app.Shutdown()
			if err != nil {
				log.Println("Error when shutdown server: ", err)
			}
		}
	}()

	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return
	}

	fmt.Println("Listerning on port: 8080")
}
