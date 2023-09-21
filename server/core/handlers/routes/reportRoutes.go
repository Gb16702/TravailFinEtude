package routes;

import "github.com/gofiber/fiber/v2"

type Report struct {
	UserID 		uint 	`json:"user_id"`
	TargetID 	uint 	`json:"target_id"`
	Reason 		string 	`json:"reason"`
}

func ReportsHandler(c *fiber.Ctx) error {

	if c.Method() != fiber.MethodPost {
		return c.Status(500).JSON(fiber.Map{
			"message": "La méthode est incorrecte",
		})
	}

	var report = new(Report);

	if err := c.BodyParser(report); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Une erreur est survenue",
		})
	}



	return nil;
}

