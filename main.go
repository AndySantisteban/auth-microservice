package main

import (
	"fmt"
	"net/http"

	"github.com/AndySantisteban/auth/authservice"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewProduction()
	defer log.Sync()
	fmt.Println("Hola, microservicio de autentificación listo!")
	err := godotenv.Load(".env")

	if err != nil {
		log.Error("Error loading .env file", zap.Error(err))
	}

	mainRouter := mux.NewRouter()

	suc := authservice.NewSignupController(log)
	sic := authservice.NewSigninController(log)

	authRouter := mainRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup", suc.SignupHandler).Methods("POST")
	authRouter.HandleFunc("/signin", sic.SigninHandler).Methods("GET")

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:4001"}))

	server := &http.Server{
		Addr:    "127.0.0.1:4001",
		Handler: ch(mainRouter),
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Error("Error Booting the Server", zap.Error(err))
	}
}
