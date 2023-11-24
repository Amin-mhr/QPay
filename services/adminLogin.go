package services

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"qpay/database"
	"qpay/models"
	"time"
)

func (a *adminInterfaceService) LoginAdmin(admin models.Admin) error {
	var adminDB models.Admin
	db := database.NewGormPostgres()
	err := db.Where("email = ?", admin.Email).First(&adminDB).Error
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(adminDB.Password), []byte(admin.Password))
	if err != nil {
		return err
	}
	return nil
}

func LoginAdminHandler(service adminInterfaceService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var admin models.Admin
		err := c.Bind(&admin)
		if err != nil {
			return err
		}
		err = service.LoginAdmin(admin)
		if err != nil {
			return err
		}
		token := GenerateAdminJWT(admin)
		setCookie(c, token)
		return c.JSON(http.StatusOK, map[string]string{
			"message": "admin logged in successfully",
		})
	}
}

func GenerateAdminJWT(admin models.Admin) string {
	claims := &JWTCustomClaims{
		Name:     admin.Name,
		Email:    admin.Email,
		Password: admin.Password,
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
