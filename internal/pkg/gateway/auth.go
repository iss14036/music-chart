package gateway

import (
	"errors"
	"fmt"
	"github.com/iss14036/music-chart/configs"
	"github.com/iss14036/music-chart/internal/pkg/constant"
	"github.com/iss14036/music-chart/internal/pkg/responsewrapper"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

//go:generate mockgen -destination=./auth_mock.go -package=gateway -source=./auth.go
type AuthItf interface {
	GetToken(userID int) (string, error)
	Middleware(next echo.HandlerFunc) echo.HandlerFunc
}

type Auth struct {
	cfg *configs.Config
}

func NewAuth(cfg *configs.Config) AuthItf {
	return &Auth{
		cfg: cfg,
	}
}

func (a *Auth) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		jwtToken := c.Request().Header.Get(constant.HeaderAuthorization)
		if jwtToken == "" {
			return responsewrapper.BadRequest(c, errors.New(""), "")
		}

		tokenString := strings.Split(jwtToken, "Bearer ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(a.cfg.AuthKey), nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return responsewrapper.Unauthorized(c, err, "")
		}

		c.Set(constant.EchoUserID, claims["user_id"])
		return next(c)
	}
}

func (a *Auth) GetToken(userID int) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(a.cfg.AuthExpiration)).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(a.cfg.AuthKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
