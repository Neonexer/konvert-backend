package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	receipts_handler "github.com/neonexer/konvert-backend/internal/core/features/receipts/handler"
	users_handler "github.com/neonexer/konvert-backend/internal/core/features/users/handler"
	users_repository "github.com/neonexer/konvert-backend/internal/core/features/users/repository"
	users_service "github.com/neonexer/konvert-backend/internal/core/features/users/service"
	core_logger "github.com/neonexer/konvert-backend/internal/core/logger"
	core_pgx_pool "github.com/neonexer/konvert-backend/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/neonexer/konvert-backend/internal/core/transport/http/middleware"
	core_http_server "github.com/neonexer/konvert-backend/internal/core/transport/http/server"
	"go.uber.org/zap"

	_ "github.com/neonexer/konvert-backend/docs/swagger"
)

// @title Konvert API
// @version 1.0
// @description API для приложения Конверт. Для управления личными финансами
// @host 127.0.0.1:5050
// @BasePath /api/v1
func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing postgres connection pool")
	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersHandler := users_handler.NewUsersHandler(usersService)

	logger.Debug("initializing feature", zap.String("feature", "receipts"))
	receiptsHandler := receipts_handler.NewReceiptsHandler()

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersionV1)

	apiVersionRouter.RegisterRoutes(usersHandler.Routes()...)
	apiVersionRouter.RegisterRoutes(receiptsHandler.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouter)
	httpServer.RegisterSwagger()

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
