package app

import (
	"auth-service/internal/handler"
	"auth-service/internal/service"
	"auth-service/internal/store/postgres"
	"auth-service/internal/store/redis"
	"auth-service/pkg/configs"
	"auth-service/pkg/databases"
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-hclog"
	"strconv"
)

type handlers interface {
	HealthCheck(c *fiber.Ctx) error
	SignUp(c *fiber.Ctx) error
}

type server struct {
	app      *fiber.App
	logger   hclog.Logger
	handlers handlers
}

func Start(conf *configs.Config) error {
	s := new(server)
	var err error

	s.logger = hclog.New(&hclog.LoggerOptions{
		JSONFormat: true,
		Level:      2,
	})

	s.app = fiber.New()

	redisDb, err := strconv.Atoi(conf.RedisDb)
	if err != nil {
		s.logger.Error("failed to parse redis number of db", "error", err)
		return err
	}

	redisConnection, err := databases.RedisConnection(conf.RedisAddress, conf.RedisPassword, redisDb)
	if err != nil {
		s.logger.Error("failed to create connection with redis", "error", err)
		return err
	}
	postgresConnection, err := databases.PostgresConnection(conf.PostgresDb)
	if err != nil {
		s.logger.Error("failed to create connection with postgres", "error", err)
		return err
	}

	postgresStore := postgres.NewStore(postgresConnection)
	redisStore := redis.NewStore(redisConnection)
	services := service.NewService(postgresStore, redisStore)

	s.handlers = handler.NewHandler(s.logger, services)

	s.router()

	err = s.app.Listen(":" + conf.Port)
	if err != nil {
		s.logger.Error("failed to listen server", "error", err)
		return err
	}

	return err
}
