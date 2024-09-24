package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type FirebaseClaim struct {
	SignInProvider string              `json:"sign_in_provider,omitempty"`
	Identities     map[string][]string `json:"identities,omitempty"`
}
type CustomClaims struct {
	Name          string           `json:"name,omitempty"`
	Picture       string           `json:"picture,omitempty"`
	UserID        string           `json:"user_id,omitempty"`
	AuthAt        *jwt.NumericDate `json:"auth_time,omitempty"`
	Email         string           `json:"email,omitempty"`
	EmailVerified bool             `json:"email_verified"`
	Firebase      *FirebaseClaim   `json:"firebase,omitempty"`

	jwt.RegisteredClaims
}

func GetJWT_Key() string {
	return os.Getenv("JWT_KEY")
}

func VerifyIDToken(idToken string) (*CustomClaims, error) {
	// Get Firebase public keys from Google
	resp, err := http.Get("https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var keys map[string]string
	if err := json.Unmarshal(body, &keys); err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Parse the ID token and validate its claims
	token, err := jwt.ParseWithClaims(idToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check token signing method etc.
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("Invalid Token : Token Header missing 'kid' (key id) field")
		}
		return jwt.ParseRSAPublicKeyFromPEM([]byte(keys[kid]))
	})

	if err != nil {
		return nil, fmt.Errorf("Token validation failed: %v", err)
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// fmt.Printf("Token is valid. Claims: %v\n", claims)
		return claims, nil
	} else {
		return nil, fmt.Errorf("Invalid token")
	}
}

// func GetEmailFromToken(claims *jwt.MapClaims) (string, error) {
// 	claims.GetAudience()
// }
