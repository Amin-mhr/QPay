package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"net/http"
	database "signup/database"
	"signup/models"
	"time"
)

type userInterface interface {
	CreateUser(user models.User) error
}

type JWTCustomClaims struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type userInterfaceService struct{}

func (u *userInterfaceService) CreateUser(user models.User) error {
	db := database.NewGormPostgres()
	err := db.Where("email = ?", user.Email).First(&user).Error
	if err == nil {
		return errors.New("email address already exists")

	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		err := db.Create(&user).Error
		if err != nil {
			return err
		}
		return nil
	}
	return err
}

func CreateHandler(service userInterfaceService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		err := c.Bind(&user)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if user.Name == "" || user.Email == "" || user.Password == "" {
			return c.JSON(http.StatusBadRequest, "please fill all fields")
		}
		err = service.CreateUser(user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		token := GenerateJWT(user)
		setCookie(c, token)
		return c.JSON(http.StatusOK, token)
	}

}

func GenerateJWT(user models.User) string {
	claims := &JWTCustomClaims{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("golang is the best"))
	if err != nil {
		return ""
	}
	return signedToken
}

func setCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "jwt-token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Minute * 5)
	c.SetCookie(cookie)
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("jwt-token")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		tokenString := cookie.Value
		token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("golang is the best"), nil
		})

		if err != nil || !token.Valid {

			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		claims, ok := token.Claims.(*JWTCustomClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		c.Set("user", claims)
		return next(c)
	}
}
func Authentication(c echo.Context) error {
	user := c.Get("user").(*JWTCustomClaims)
	return c.String(http.StatusOK, "Welcome "+user.Name)
}

func SignUpRoutes(server *echo.Echo) {
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.POST("/signup", CreateHandler(userInterfaceService{}))
	server.GET("/home", Authentication, authMiddleware)
}
