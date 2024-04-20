package internal

import (
	"context"
	"currency/internal/config"
	"currency/internal/database"
	v1 "currency/internal/http/v1"
	"currency/internal/integration"
	"currency/internal/logging"
	"currency/internal/repository"
	"currency/internal/server"
	"currency/internal/service"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	srv *server.Server
	cfg *config.Configs
}

func Init() *App {
	if err := logging.InitDefaultLogger(); err != nil {
		zap.S().Fatalf("Can't initialize logger: %v", err)
	}

	zap.S().Info("Starting currency service at ", time.Now().Local())
	zap.S().Info("Loading config...")

	if err := config.LoadLocalConfig(config.Conf); err != nil {
		zap.S().Fatalf("Config load error! %v", err)
	}

	cfg := config.Conf

	db := database.InitializeDB(cfg.DB)
	itg := integration.NewIntegration(&cfg.NationalBank)
	repo := repository.NewRepository(cfg, db)
	svc := service.NewService(repo, itg)
	ctrl := v1.NewController(svc)
	ctrl.InitRoutes()

	srv := server.NewServer(ctrl.StartRoutes())

	return &App{
		srv: srv,
		cfg: cfg,
	}
}

// Run Запуск приложения с Graceful Shutdown
func (a *App) Run() {
	// Канал для принятия сигналов
	quit := make(chan os.Signal, 1)

	// Принимаем любые внешние сигналы. Пример: Interrupt
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	// Запуск сервера в горутине
	go func() {
		zap.S().Infof("Server is started at port: %s", a.cfg.Port)
		a.srv.Run()
	}()

	// Блокировка основного потока (горутины) до появления любой ошибки или сигнала
	select {
	// case для внешних сигналов
	case sig := <-quit:
		zap.S().Infof("app: signal accepted: %s", sig.String())
	// case для ошибок сервера
	case err := <-a.srv.Notify():
		zap.S().Infof("app: signal accepted: %v", err)
	}

	// Выключение сервера, если произошло что-либо из перечисленного
	if err := a.srv.Shutdown(context.Background()); err != nil {
		zap.S().Errorf("error while server shutting down: %v", err)
	}
}
