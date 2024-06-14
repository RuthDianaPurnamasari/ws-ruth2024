package main

import (
	"log"

	"github.com/RuthDianaPurnamasari/ws-ruth2024/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"


	"github.com/RuthDianaPurnamasari/ws-ruth2024/url"

	"github.com/gofiber/fiber/v2"
	_ "github.com/RuthDianaPurnamasari/ws-ruth2024/docs"

)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/RuthDianaPurnamasari
// @contact.email 714220042@std.ulbi.ac.id

// @host ws-ruth2024-458a1e70d2e1.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
