package config

import (
	"E-Commerce/models/dto"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

func ConnectToDB(in dto.ConfigData, logger zerolog.Logger) (*sql.DB, error) {
	// connect to database
	logger.Info().Msg("Try connect to database...")

	// init database connection string
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", in.DbConfig.User, in.DbConfig.Pass, in.DbConfig.DbPort, in.DbConfig.Host, in.DbConfig.Database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to open database connections")
		return nil, err
	}

	// test ping to database
	if err = db.Ping(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to ping database")
		return nil, err
	}

	logger.Info().Msg("Database connected successfully")
	return db, err
}
