package core

import (
	"host/core/handlers"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

/**
*	This function handles the server start
*
*	@param Url string
*	@returns void
*/
func HandleServerStart(Url string) {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	});
	app.Use(cors.New(), logger.New());

	handlers.Router(app);

	app.Listen(Url);
}
