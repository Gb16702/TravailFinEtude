package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

/**
*	This function handles the server start
*
*	@param string Url
*	@returns void
*/
func HandleServerStart(Url string) {
	app := fiber.New();
	app.Use(cors.New(), logger.New());

	app.Listen(Url);
}
