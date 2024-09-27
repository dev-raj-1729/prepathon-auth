package handlers

import (
	"prepathon-auth/controllers"
	"prepathon-auth/models"
	"prepathon-auth/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateUserWithFirebaseToken(c *fiber.Ctx) error {
	idToken := c.Get("firebase_token", "")

	claims, err := utils.VerifyIDToken(idToken)

	if err != nil {
		return err
	}
	user := models.User{
		Email:    claims.Email,
		Name:     claims.Name,
		PhotoURL: claims.Picture,
	}
	if err := controllers.CreateUser(h.MongoClient, &user); err != nil {
		return err
	}

	return c.JSON(user)

}

func (h *Handler) FindUserWithFirebaseToken(c *fiber.Ctx) error {
	firebaseIdToken := c.Get("firebase_token", "")

	claims, err := utils.VerifyIDToken(firebaseIdToken)

	if err != nil {
		return err
	}

	user, err := controllers.FindUserByEmail(h.MongoClient, claims.Email)

	if err != nil {
		return err
	}

	return c.JSON(user)
}
