package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt"
)

var SECRET_KEY = "keepTheSecret"

type MyClaims struct {
	User_id uuid.UUID `json:"user_id"`
	jwt.StandardClaims
	Username     string `json:"username"`
	Access_level string `json:"access_level"`
}

func GenerateToken(uuid uuid.UUID, username, access_level string) (string, error) {

	claims := MyClaims{
		User_id: uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
		Username:     username,
		Access_level: access_level,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	responseToken, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", err
	}

	return responseToken, nil
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(SECRET_KEY), nil
	})

	validationExpiredJWT := token.Claims.Valid()
	if validationExpiredJWT != nil {
		return nil, validationExpiredJWT
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
