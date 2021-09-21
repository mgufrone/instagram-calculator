package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mgufrone/insta-calculator/server/handlers"
	"github.com/mgufrone/insta-calculator/server/helpers"
	"github.com/mgufrone/insta-calculator/server/service/instagram"
	"go.uber.org/fx"
	"log"
	"net/http"
	"os"
)

func server() (*gin.Engine, *http.Server) {
	eng := gin.New()
	eng.Use(cors.Default())
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: eng,
	}
	return eng, srv
}
func registerRoutes(srv *gin.Engine, ig *handlers.Instagram) {
	srv.POST("/profile", ig.GetProfile)
}
func startServer(srv *http.Server, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("start server at ", os.Getenv("PORT"))
			return srv.ListenAndServe()
		},
		OnStop: func(ctx context.Context) error {
			log.Println("closing connection")
			return srv.Close()
		},
	})
}
func app() *fx.App {
	godotenv.Load()
	return fx.New(
		fx.Provide(
			func() *http.Client {
				cli := &http.Client{
					CheckRedirect: func(req *http.Request, via []*http.Request) error {
						return http.ErrUseLastResponse
					},
					Transport: helpers.NewRoundTripper(http.DefaultTransport),
				}
				return cli
			},
			instagram.New,
			handlers.NewHandler,
			server,
		),
		fx.Invoke(
			registerRoutes,
			startServer,
		),
	)
}
func main() {
	a := app()
	if err := a.Start(context.Background()); err != nil {
		panic(err)
	}
	<-a.Done()
}
