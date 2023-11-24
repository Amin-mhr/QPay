package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"qpay/database"
	"qpay/models"
	"time"
)

type UserInterface interface {
	RegisterUser(user models.User) error
	LoginUser(user models.User) error
}

type JWTCustomClaims struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type UserInterfaceService struct{}

func (u *UserInterfaceService) RegisterUser(user models.User) error {
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

func (u *UserInterfaceService) LoginUser(user models.User) error {
	db := database.NewGormPostgres()
	var userDB models.User
	err := db.Where("email = ?", user.Email).First(&userDB).Error
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
	if err != nil {
		return err
	}
	return nil
}

func encryptPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func RegisterHandler(service UserInterfaceService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		err := c.Bind(&user)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if user.Name == "" || user.Email == "" || user.Password == "" {
			return c.JSON(http.StatusBadRequest, "please fill all fields")
		}
		user.Password = encryptPassword(user.Password)
		err = service.RegisterUser(user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		token := GenerateJWT(user)
		setCookie(c, token)
		return c.JSON(http.StatusOK, token)
	}

}

func LoginHandler(service UserInterfaceService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		err := c.Bind(&user)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if user.Email == "" || user.Password == "" {
			return c.JSON(http.StatusBadRequest, "please fill all fields")
		}
		err = service.LoginUser(user)
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

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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
	server.POST("/signup", RegisterHandler(UserInterfaceService{}))
	server.POST("/login", LoginHandler(UserInterfaceService{}))
	server.GET("/home", Authentication, AuthMiddleware)
}
