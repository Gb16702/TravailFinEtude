package routes

import (
	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	"host/database"
	"host/database/models"
	"host/utils"
)

// This function handles the auth routes
//
// @param *fiber.Ctx c
// @returns error
//func AuthHandler(c *fiber.Ctx) error {
//	registerErr := register(c)
//	if registerErr != nil {
//		log.Fatal(
//			"Erreur lors de l'appel de la fonction register: ", registerErr,
//		)
//	}
//
//	loginErr := login(c)
//	if loginErr != nil {
//		log.Fatal(
//			"Erreur lors de l'appel de la fonction login:", loginErr,
//		)
//	}
//
//	return nil
//}

// This function handles the register route
//
// @param *fiber.Ctx c
// @returns error
func Register(c *fiber.Ctx) error {
	if c.Method() != fiber.MethodPost {
		return c.Status(500).JSON(fiber.Map{
			"message": "La méthode est incorrecte",
		})
	}

	type Register_struct struct {
		FirstName       string `json:"firstname"`
		LastName        string `json:"lastname"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		PasswordConfirm string `json:"passwordconfirm"`
		PhoneNumber     string `json:"phone_number"`
	}

	Register := Register_struct{}
	if err := c.BodyParser(&Register); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Une erreur est survenue" + err.Error(),
		})
	}

	requiredFields := []string{"firstname", "lastname", "email", "password", "passwordconfirm", "phone_number"}
	for _, field := range requiredFields {
		switch field {
		case "firstname":
			if utils.IsNilOrEmpty(Register.FirstName) {
				return c.Status(500).JSON(fiber.Map{
					"message": "Le prénom est requis",
				})
			}

		case "lastname":
			if utils.IsNilOrEmpty(Register.LastName) {
				return c.Status(500).JSON(fiber.Map{
					"message": "Le nom est requis",
				})
			}

		case "email":
			if utils.IsNilOrEmpty(Register.Email) {
				return c.Status(500).JSON(fiber.Map{
					"message": "L'email est requis",
				})
			}

			if _, err := utils.IsValidEmail(Register.Email); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			if err := utils.IsEmailAlreadyExisting(Register.Email); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

		case "password":
			if utils.IsNilOrEmpty(Register.Password) {
				return c.Status(500).JSON(fiber.Map{
					"message": "Le mot de passe est requis",
				})
			}

			if _, err := utils.IsValidPassword(Register.Password); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

		case "passwordconfirm":
			if utils.IsNilOrEmpty(Register.PasswordConfirm) {
				return c.Status(500).JSON(fiber.Map{
					"message": "La confirmation du mot de passe est requise",
				})
			}

			if _, err := utils.ArePasswordMatching(Register.Password, Register.PasswordConfirm); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

		case "phone_number":
			if utils.IsNilOrEmpty(Register.PhoneNumber) {
				return c.Status(500).JSON(fiber.Map{
					"message": "Le numéro de téléphone est requis",
				})
			}
			if isValid, err := utils.IsValidPhoneNumber(Register.PhoneNumber); !isValid {
				return c.Status(500).JSON(fiber.Map{
					"message": "Le numéro de téléphone est invalide" + err,
				})
			}
		}
	}

	hash, err := utils.HashPassword(Register.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Une erreur est survenue",
		})
	}

	user := models.User{
		FirstName:   Register.FirstName,
		LastName:    Register.LastName,
		PhoneNumber: Register.PhoneNumber,
		UserLogin: models.UserLogin{
			Email:    Register.Email,
			Password: hash,
		},
	}
	if result := database.DB.Create(&user); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "L'utilisateur n'a pas pu être enregistré" + result.Error.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":  "L'utilisateur a été enregistré",
		"register": Register,
		"hash":     hash,
	})
}

func login(c *fiber.Ctx) error {
	if c.Method() != fiber.MethodPost {
		return c.Status(500).JSON(fiber.Map{
			"message": "La méthode est incorrecte",
		})
	}

	var login = models.UserLogin{}
	if err := c.BodyParser(&login); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "La requête est incorrecte",
		})
	}

	if utils.IsNilOrEmpty(&login.Email) {
		return c.Status(500).JSON(fiber.Map{
			"message": "L'email est requise",
		})
	}

	if utils.IsNilOrEmpty(&login.Password) {
		return c.Status(500).JSON(fiber.Map{
			"message": "Le mot de passe est requis",
		})
	}

	user := models.User{}
	result := database.DB.Where("email = ?", login.Email).First(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Identifiants invalides 2",
		})
	}

	match, _, err := argon2id.CheckHash(login.Password, user.UserLogin.Password)
	if err != nil || !match {

		return c.Status(500).JSON(fiber.Map{
			"message": "Identifiants invalides 3",
		})
	}

	session, err := utils.GenerateToken(user.UserLogin.Email, user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "La session n'a pas pu être créée",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Authentification réussie",
		"session": session,
	})
}
