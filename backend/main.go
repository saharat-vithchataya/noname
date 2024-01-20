package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/saharat-vithchataya/noname/handlers"
	"github.com/saharat-vithchataya/noname/repository"
	"github.com/saharat-vithchataya/noname/services"
)

func main() {
	// lo
	dsn := "postgres://postgres:example@localhost:5432/sinestrea?sslmode=disable"
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	accountRepository := repository.NewAccountRepositoryDB(db)
	accountService := services.NewAccountService(accountRepository)
	accountHandler := handlers.NewAccountHandler(accountService)

	authService := services.NewAuthService(accountRepository)
	authHandler := handlers.NewAuthenticationHandler(authService)

	app := fiber.New()
	app.Get("/:account_id", accountHandler.GetAccount)
	app.Post("/", accountHandler.OpenNewAccount)
	app.Post("/login", authHandler.Login)
	app.Listen(":8000")
}
