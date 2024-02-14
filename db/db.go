package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"
)

type Db struct {
	Pool *pgxpool.Pool
}

type Config struct {
	username string
	password string
	database string
	host     string
}

var (
	instance *Db
	once     sync.Once
)

func CreateDb() *Db {

	once.Do(func() {
		config := Config{username: "postgres", password: "postgres", host: "localhost", database: "url_shorter"}
		pool, err := pgxpool.New(context.Background(), fmt.Sprintf("postgresql://%s:%s@%s/%s", config.username, config.password, config.host, config.database))

		if err != nil {
			log.Fatal("Database connection failed")
		}

		instance = &Db{Pool: pool}
	})

	return instance
}
