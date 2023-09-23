package handlers

import (
	"host/core/handlers/routes"

	"github.com/gofiber/fiber/v2"
)

/**
*	This function handles the server routes
*
*	@param App *fiber.App
*	@returns void
 */
func Router(App *fiber.App) {
	apiGroup := App.Group("/api");
		authGroup := apiGroup.Group("/auth");
			authGroup.Post("/register", routes.AuthHandler);

}
