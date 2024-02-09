package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go-clean-architecture/config"
	"go-clean-architecture/internal/admin"
	adminhttp "go-clean-architecture/internal/admin/delivery/http"
	usecase3 "go-clean-architecture/internal/admin/usecase"
	"go-clean-architecture/internal/auth"
	authhttp "go-clean-architecture/internal/auth/delivery/http"
	authpostgres "go-clean-architecture/internal/auth/repository/postgres"
	"go-clean-architecture/internal/auth/usecase"
	"go-clean-architecture/internal/news"
	newshttp "go-clean-architecture/internal/news/delivery/http"
	newspostgres "go-clean-architecture/internal/news/repository/postgres"
	usecase2 "go-clean-architecture/internal/news/usecase"
	"go-clean-architecture/pkg/utils/logger"
	"go-clean-architecture/pkg/utils/usersession"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server

	authUC  auth.UseCase
	newsUC  news.UseCase
	adminUC admin.UseCase
}

func NewApp(appConfig *config.Config) *App {
	initSession(&appConfig.Session)
	db, err := initDB(&appConfig.DB)
	if err != nil {
		logger.GetLogger().Fatalf("%s", err)
	}

	roleRepo := authpostgres.NewRoleRepository(db)
	userRepo := authpostgres.NewUserRepository(db)
	newsRepo := newspostgres.NewNewsRepository(db)

	return &App{
		authUC: usecase.NewAuthUseCase(
			userRepo,
			roleRepo,
		),
		newsUC: usecase2.NewNewsUseCase(
			newsRepo,
		),
		adminUC: usecase3.NewAdminUseCase(
			userRepo,
			roleRepo,
			newsRepo,
		),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	router.LoadHTMLGlob("ui/html/*")
	fileServer := http.FileServer(http.Dir("ui/static/"))
	router.GET("/static/*filepath", func(c *gin.Context) {
		fileServer.ServeHTTP(c.Writer, c.Request)
	})

	adminMiddleware := adminhttp.NewAdminMiddleware(a.adminUC)
	authMiddleware := authhttp.NewAuthMiddleware(a.authUC)

	api := router.Group("/api", authMiddleware)
	adminApi := router.Group("/admin", adminMiddleware)

	authhttp.AuthHTTPEndpoints(router, a.authUC)
	newshttp.NewsHTTPEndpoints(api, a.newsUC)
	adminhttp.AdminHTTPEndpoints(adminApi, a.adminUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			logger.GetLogger().Fatalf("Failed to listen and serve:: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB(dbConfig *config.DBConf) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initSession(sessionConfig *config.Session) {
	usersession.NewUserSessionStore(sessionConfig)
}
