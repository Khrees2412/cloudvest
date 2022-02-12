package auth

import (
	// "fmt"

	db "risevest/database"
	"risevest/models"
	"risevest/utils"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Data struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	user := new(models.User)  // from database
	input := new(models.User) //from frontend

	if err := c.BodyParser(input); err != nil {
		log.Error(err)
		return c.JSON(fiber.Map{"error": true, "message": "Please review your input"})
	}
	// &models.User{Email: input.Email}
	if res := db.DB.Where(
		"email = ?", input.Email).Find(&user); res.RowsAffected <= 0 {
		return c.JSON(fiber.Map{"error": true, "message": "Invalid Credentials."})
	}

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.Error(err)
		return c.JSON(fiber.Map{"success": false, "message": "Invalid Credentials."})
	}

	// setting up the authorization cookies
	accessToken, refreshToken := utils.GenerateTokens(user.UUID.String())
	accessCookie, refreshCookie := utils.GetAuthCookies(accessToken, refreshToken)
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
