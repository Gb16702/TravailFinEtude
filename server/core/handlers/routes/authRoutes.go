package routes

import (
	"host/database"
	"host/database/models"
	"host/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

/**
*	This function handles the auth routes
*
*	@param *fiber.Ctx c
*	@returns error
 */
func AuthHandler(c *fiber.Ctx) error {
	err := register(c);

	if err != nil {
		log.Fatal(err);
	}

	return nil;
}

/**
*	This function handles the register route
*
*	@param *fiber.Ctx c
*	@returns error
*/
func register(c *fiber.Ctx) error {
	type register_struct struct {
		FirstName 			string 	`json:"firstname"`
		LastName 			string 	`json:"lastname"`
		Email 				string 	`json:"email"`
		Password 			string 	`json:"password"`
		PasswordConfirm		string 	`json:"passwordconfirm"`
	}

	if c.Method() != fiber.MethodPost {
		return c.Status(500).JSON(fiber.Map{
			"message": "La méthode est incorrecte",
		});
	}

	register := register_struct{};
	if err := c.BodyParser(&register); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Une erreur est survenue" + err.Error(),
		})
	}

	requiredFields := []string{"firstname", "lastname", "email", "password", "passwordconfirm"}
	for _, field := range requiredFields {
		switch field {
		case "firstname":
			if utils.IsNilOrEmpty(register.FirstName) {
				return c.Status(500).JSON(fiber.Map{
					"message": "Le prénom est requis",
				})
			}

		case "lastname":
			if utils.IsNilOrEmpty(register.LastName) {
				return c.Status(500).JSON(fiber.Map{
					"message": "Le nom est requis",
				})
			}

		case "email":
			if utils.IsNilOrEmpty(register.Email) {
				return c.Status(500).JSON(fiber.Map{
					"message": "L'email est requis",
				})
			}

			if _, err := utils.IsValidEmail(register.Email); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			isExistingEmail := database.DB.Where("email = ?", register.Email).First(&models.User{})
			if isExistingEmail.RowsAffected != 0 {
				return c.Status(500).JSON(fiber.Map{
					"message": "L'email est déjà utilisé",
				})
			}

		case "password":
			if utils.IsNilOrEmpty(register.Password) {
				return c.Status(500).JSON(fiber.Map{
					"message": "Le mot de passe est requis",
				})
			}

			if _, err := utils.IsValidPassword(register.Password); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

		case "passwordconfirm":
			if utils.IsNilOrEmpty(register.PasswordConfirm) {
				return c.Status(500).JSON(fiber.Map{
					"message": "La confirmation du mot de passe est requise",
				})
			}

			if _, err := utils.ArePasswordMatching(register.Password, register.PasswordConfirm); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}
		}
	}

	hash, err := utils.HashPassword(register.Password);
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Une erreur est survenue",
		})
	}

	user := models.User{
		FirstName: register.FirstName,
		LastName: register.LastName,
		UserLogin: models.UserLogin{
			Email: register.Email,
			Password: hash,
		},
	}
	if result := database.DB.Create(&user); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "L'utilisateur n'a pas pu être enregistré" + result.Error.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "L'utilisateur a été enregistré",
		"register": register,
		"hash": hash,
	})
}