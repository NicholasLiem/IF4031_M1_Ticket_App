package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DAO interface {
	NewEventQuery() EventQuery
	NewSeatQuery() SeatQuery
}

type dao struct {
	pgdb  *gorm.DB
	redis *redis.Client
}

var DisableLogger = logger.New(
	nil, // Use the default logger output, which is discarded
	logger.Config{
		LogLevel: logger.Silent,
	},
)

func NewDAO(db *gorm.DB, rc *redis.Client) DAO {
	return &dao{
		pgdb:  db,
		redis: rc,
	}
}

func SetupDB() *gorm.DB {
	var dbHost = os.Getenv("DB_HOST")
	var dbName = os.Getenv("DB_NAME")
	var dbUsername = os.Getenv("DB_USERNAME")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbPort = os.Getenv("DB_PORT")
	var envMode = os.Getenv("ENVIRONMENT")

	var dsn string
	if envMode == "DEV" {
		dsn = fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", dbUsername, dbPassword, dbName, dbPort)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUsername, dbPassword, dbName, dbPort)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: DisableLogger,
	})

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get DB instance: " + err.Error())
	}

	log.Println("Successfully connected to DB")

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}

func SetupRedis(ctx context.Context) *redis.Client {
	redisDB, err := strconv.ParseInt(os.Getenv("REDIS_DB"), 10, 64)
	if err != nil {
		panic("Failed to parse REDIS_DB: " + err.Error())
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       int(redisDB),
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	log.Println("Successfully connected to Redis")

	return client
}

func (d *dao) NewEventQuery() EventQuery {
	return NewEventQuery(d.pgdb, d.redis)
}

func (d *dao) NewSeatQuery() SeatQuery {
	return NewSeatQuery(d.pgdb)
}
