package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/caarlos0/env"
	_ "github.com/lib/pq"
)

type Config struct {
	Port       string `env:"APP_CONTAINER_PORT"`
	DBHost     string `env:"POSTGRES_HOST"`
	DBPort     string `env:"POSTGRES_PORT"`
	DBUser     string `env:"POSTGRES_USER"`
	DBPassword string `env:"POSTGRES_PASSWORD"`
	DBName     string `env:"POSTGRES_DB"`
	DBSSLMODE  string `env:"POSTGRES_SSLMODE"`
}

func getConfig() (*Config, error) {
	config := &Config{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}
	return config, nil
}

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMODE)
	dbHandle, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	ctx := context.Background()
	conn, err := dbHandle.Conn(ctx)
	if err != nil {
		log.Fatal("DB開かねえな、なんで: %v", err)
	}

	r, err := conn.ExecContext(ctx, "INSERT INTO users(name) VALUES($1)", "uehara")
	if err != nil {
		log.Println("ふざけんな")
	}
	fmt.Println(r.RowsAffected())

	//fmt.Println("Golang postgres boilerplate managed by taskfile.")
}
