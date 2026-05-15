package utils

import (
	"backend/dto"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID, username, role string) (string, error) {
	payload := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("error konfigurasi server")
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (*dto.TokenPayload, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return nil, errors.New("error konfigurasi server")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("klaim token tidak valid")
	}

	userID, ok := claimString(claims, "user_id")
	if !ok {
		return nil, errors.New("klaim user_id tidak valid")
	}
	username, ok := claimString(claims, "username")
	if !ok {
		return nil, errors.New("klaim username tidak valid")
	}
	role, ok := claimString(claims, "role")
	if !ok {
		return nil, errors.New("klaim role tidak valid")
	}

	payload := &dto.TokenPayload{
		UserID:   userID,
		Username: username,
		Role:     role,
	}

	return payload, nil
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		// Try to get token from Authorization header first
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				tokenString = parts[1]
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Format token tidak valid"})
				c.Abort()
				return
			}
		} else {
			// Fallback to query parameter (for WebSocket)
			// WebSocket upgrade doesn't allow custom headers per HTTP spec
			tokenString = c.Query("token")
			if tokenString == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Tidak Terautentikasi"})
				c.Abort()
				return
			}
		}

		secretKey := os.Getenv("JWT_SECRET")
		if secretKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Kesalahan konfigurasi server"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Klaim token tidak valid"})
			c.Abort()
			return
		}

		userID, ok := claimString(claims, "user_id")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Klaim user_id tidak valid"})
			c.Abort()
			return
		}
		username, ok := claimString(claims, "username")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Klaim username tidak valid"})
			c.Abort()
			return
		}
		role, ok := claimString(claims, "role")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Klaim role tidak valid"})
			c.Abort()
			return
		}

		payload := &dto.TokenPayload{
			UserID:   userID,
			Username: username,
			Role:     role,
		}

		c.Set("userID", payload.UserID)
		c.Set("username", payload.Username)
		c.Set("userRole", payload.Role)
		c.Next()
	}
}

func claimString(claims jwt.MapClaims, key string) (string, bool) {
	value, ok := claims[key].(string)
	return value, ok && value != ""
}
