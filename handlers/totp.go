package handlers

import (
	"fmt"
	"prepathon-auth/controllers"
	"prepathon-auth/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
)

func (h *Handler) Enable2FA(c *fiber.Ctx) error {

	firebaseToken := c.Get("firebase_token", "")
	claims, err := utils.VerifyIDToken(firebaseToken)

	if err != nil {
		return err
	}

	user, err := controllers.FindUserByEmail(h.MongoClient, claims.Email)
	if err != nil {
		return err
	}

	secret := totp.GenerateOpts{
		Issuer:      "TZoO Prepathon",
		AccountName: claims.Email,
	}

	key, err := totp.Generate(secret)

	if err != nil {
		return err
	}

	if err := controllers.CreateSecret(h.MongoClient, key.Secret(), user.ID); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"otpauth_url": key.URL(),
	})
}

func (h *Handler) Verify2FA(c *fiber.Ctx) error {

	firebaseToken := c.Get("firebase_token", "")
	totpToken := c.Get("totp_token", "")
	fmt.Println(totpToken)
	claims, err := utils.VerifyIDToken(firebaseToken)

	if err != nil {
		return err
	}

	user, err := controllers.FindUserByEmail(h.MongoClient, claims.Email)

	if err != nil {
		return err
	}

	totpSecret, err := controllers.FindSecretByUserId(h.MongoClient, user.ID)
	if err != nil {
		return err
	}
	isValid := totp.Validate(totpToken, totpSecret.Secret)
	fmt.Println(totp.GenerateCode(totpSecret.Secret, time.Now()))
	if !isValid {
		return fiber.ErrForbidden
	}

	expirationTime := time.Now().Add(time.Hour * 24 * 30)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     expirationTime.Unix(),
	})
	jwtKey := utils.GetJWT_Key()
	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"access_token": tokenString,
	})
}
