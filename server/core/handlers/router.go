package handlers

import (
	"github.com/gofiber/fiber/v2"
	"host/core/handlers/routes"
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
