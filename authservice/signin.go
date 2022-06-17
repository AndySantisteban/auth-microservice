package authservice

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/AndySantisteban/auth/authservice/data"
	"github.com/AndySantisteban/auth/authservice/jwt"
	"go.uber.org/zap"
)

type SigningController struct {
	logger *zap.Logger
}

func NewSigninController(logger *zap.Logger) *SigningController {
	return &SigningController{
		logger: logger,
	}
}

func getSignedToken() (string, error) {
	claimsMap := map[string]string{
		"aud": "frontend.knowsearch.ml",
		"iss": "knowsearch.ml",
		"exp": fmt.Sprint(time.Now().Add(time.Minute * 1).Unix()),
	}

	secret := jwt.GetSecret()
	if secret == "" {
		return "", errors.New("contraseña secreta de JWT no encontrada")
	}

	header := "HS256"
	tokenString, err := jwt.GenerateToken(header, claimsMap, secret)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func validateUser(email string, passwordHash string) (bool, error) {
	usr, exists := data.GetUserObject(email)
	if !exists {
		return false, errors.New("usuario no existe")
	}
	passwordCheck := usr.ValidatePasswordHash(passwordHash)

	if !passwordCheck {
		return false, nil
	}
	return true, nil
}
func (ctrl *SigningController) SigninHandler(rw http.ResponseWriter, r *http.Request) {

	if _, ok := r.Header["Email"]; !ok {
		ctrl.logger.Warn("ingrese el correo electrónico")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Email no encontrado"))
		return
	}
	if _, ok := r.Header["Passwordhash"]; !ok {
		ctrl.logger.Warn("Contraseña no encontrada")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Contraseña no encontrada"))
		return
	}
	valid, err := validateUser(r.Header["Email"][0], r.Header["Passwordhash"][0])
	if err != nil {
		ctrl.logger.Warn("Usuario no existente", zap.String("email", r.Header["Email"][0]))
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Usuario no existente"))
		return
	}

	if !valid {
		ctrl.logger.Warn("Contraseña incorrecta", zap.String("email", r.Header["Email"][0]))
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Contraseña incorrecta"))
		return
	}
	tokenString, err := getSignedToken()
	if err != nil {
		ctrl.logger.Error("No puedes iniciar sesion ahorita", zap.Error(err))
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Internal Server Error"))
		return
	}
	ctrl.logger.Info("Token", zap.String("token", tokenString), zap.String("email", r.Header["Email"][0]))

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(tokenString))
}
