package controller

import (
	"my-part/database"
	"my-part/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c echo.Context) error {
	var data map[string]string

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request"})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to hash password"})
	}

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: string(password),
	}

	if err := database.NewGormPostgres().Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error"})
	}

	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	var data map[string]string

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request"})
	}

	var user models.User

	if err := database.NewGormPostgres().Where("email = ?", data["email"]).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error"})
	}

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "user not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "incorrect password"})
	}

	// Use jwt.TimeFunc() to get the current time
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.NewTime(time.Hour.Hours()),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "could not generate token"})
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	http.SetCookie(c.Response(), cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}

func User(c echo.Context) error {
	cookie, err := c.Cookie("jwt")

	if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized"})
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error){
		return []byte(SecretKey), nil
	})

	//این جا مشکل داره، میگه که توکنت اکسپایر شده
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized"})
	}

	if !token.Valid {
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid token"})
    }

    claims, ok := token.Claims.(*jwt.StandardClaims)
    if !ok {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error"})
    }
	return c.JSON(http.StatusOK, claims)
}
