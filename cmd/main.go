package main

import (
	"net/http"
	"stepik_1/configs"
	"stepik_1/internal/auth"
	"stepik_1/internal/link"
	"stepik_1/internal/user"
	"stepik_1/pkg/db"
	"stepik_1/pkg/middleware"
)

func main() {
	conf := configs.LoadCongig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)

	//handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepo: linkRepository,
	})

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}
	println("http server is running")
	err := server.ListenAndServe()
	if err != nil {
		println("server faild: ", err.Error())
	}
}
