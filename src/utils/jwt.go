package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Clims struct {
	UserId uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

var secret = []byte("hh")

func GetToken(userid uuid.UUID) (string, error) {
	clims := &Clims{
		userid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 7 * 24)),
			Issuer:    "hh",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clims)
	return token.SignedString(secret)
}

// ada9823b-02e5-4473-931e-1b5eedffac60

func ParseToken(tokenString string) (*Clims, error) {
	cm := new(Clims)
	token, err := jwt.ParseWithClaims(tokenString, cm, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return cm, nil
	}
	return nil, errors.New("invalid token")
}

func ValidRequestToken(c *gin.Context) {
	Token := c.Request.Header["H-Token"]
	if len(Token) == 0 {
		c.JSON(200, BackMessageResp(200, "token为空"))
		c.Abort()
		return
	}
	cm, err := ParseToken(Token[0])
	if err != nil {
		c.JSON(200, BackMessageResp(201, "invalid token"))
		c.Abort()
		return
	}
	if cm != nil {
		c.Set("userid", cm.UserId)
		c.Next()
	}
}
