package middleware

import (
	"net/http"

	"github.com/AndySantisteban/auth/authservice/jwt"
	"go.uber.org/zap"
)

type TokenMiddleware struct {
	logger *zap.Logger
}

func NewTokenMiddleware(logger *zap.Logger) *TokenMiddleware {
	return &TokenMiddleware{
		logger: logger,
	}
}

func (ctrl *TokenMiddleware) TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if _, ok := r.Header["Token"]; !ok {
			ctrl.logger.Warn("Token no encontrado en el header")
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Token Premise Denegado"))
			return
		}
		token := r.Header["Token"][0]

		secret := jwt.GetSecret()
		if secret == "" {
			ctrl.logger.Error("Contraseña secreta de JWT no encontrada")
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Internal Server Error"))
			return
		}

		check, err := jwt.ValidateToken(token, secret)
		if err != nil {
			ctrl.logger.Error("Token no valido", zap.String("token", token))
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Token no valido"))
			return
		}
		if !check {
			ctrl.logger.Warn("Token no valido", zap.String("token", token))
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Token no valido"))
			return
		}
		next.ServeHTTP(rw, r)
	})
}
