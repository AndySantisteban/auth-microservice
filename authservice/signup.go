package authservice

import (
	"net/http"

	"github.com/AndySantisteban/auth/authservice/data"
	"go.uber.org/zap"
)

type SignupController struct {
	logger *zap.Logger
}

func NewSignupController(logger *zap.Logger) *SignupController {
	return &SignupController{
		logger: logger,
	}
}

func (ctrl *SignupController) SignupHandler(rw http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Email"]; !ok {
		ctrl.logger.Warn("Email no enviado")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Nombre de usuario no enviado"))
		return
	}
	if _, ok := r.Header["Username"]; !ok {
		ctrl.logger.Warn("Nombre de usuario no enviado")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Nombre de usuario no enviado"))
		return
	}
	if _, ok := r.Header["Passwordhash"]; !ok {
		ctrl.logger.Warn("Contraseña no enviado")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Contraseña no enviado"))
		return
	}
	if _, ok := r.Header["Fullname"]; !ok {
		ctrl.logger.Warn("Nombre completo no enviado")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Nombre completo no enviado"))
		return
	}

	check := data.AddUserObject(r.Header["Email"][0], r.Header["Username"][0], r.Header["Passwordhash"][0],
		r.Header["Fullname"][0], 0)
	if !check {
		ctrl.logger.Warn("Usuario existente", zap.String("email", r.Header["Email"][0]), zap.String("username", r.Header["Username"][0]))
		rw.WriteHeader(http.StatusConflict)
		rw.Write([]byte("Usuario existente"))
		return
	}
	ctrl.logger.Info("User Creado", zap.String("email", r.Header["Email"][0]), zap.String("username", r.Header["Username"][0]))
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User Creado"))
}
