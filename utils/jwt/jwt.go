package jwt

import (
	"fmt"
	"os"
	"strings"
	"time"

	"todo/utils/encryption"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId int64, days int) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	iss := os.Getenv("APP_DOMAIN")
	sub := userId
	iat := int32(time.Now().Unix())
	subIat := fmt.Sprintf("%d%d", sub, iat)
	jti := encryption.MD5Hash(subIat)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": iss,
		"sub": sub,
		"iat": iat,
		"jti": jti,
		"exp": time.Now().Add(time.Hour * 24 * time.Duration(days)).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func ParseToken(authorizationHeader string) map[string]interface{} {
	if authorizationHeader == "" {
		errRes := map[string]interface{}{
			"status":      "errors",
			"status_code": "empty_token",
			"message":     "An authorization header is required",
		}

		return errRes
	}

	bearerToken := strings.Split(authorizationHeader, " ")
	if len(bearerToken) != 2 {
		errRes := map[string]interface{}{
			"status":      "errors",
			"status_code": "invalid_bearer_token",
			"message":     "Invalid Bearer Token",
		}

		return errRes
	}

	token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		errRes := map[string]interface{}{
			"status":      "errors",
			"status_code": "token_expired",
			"message":     err.Error(),
		}

		return errRes
	}

	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			errRes := map[string]interface{}{
				"status":      "errors",
				"status_code": "invalid_jwt",
				"message":     "Invalid JWT Token",
			}

			return errRes
		}

		userIdFloat64 := claims["sub"].(float64)
		userIdInt64 := int64(userIdFloat64)

		results := map[string]interface{}{
			"status":  "success",
			"message": "Success",
			"user_id": userIdInt64,
		}

		return results
	}

	errRes := map[string]interface{}{
		"status":      "errors",
		"status_code": "invalid_authorization_token",
		"message":     "Invalid authorization token",
	}

	return errRes
}
