package app

import (
	"E-Commerce/config"
	"E-Commerce/models/dto"
	"E-Commerce/pkg/middleware"
	"E-Commerce/pkg/utils"
	"E-Commerce/router"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func initEnv() (dto.ConfigData, error) {
	// load env data
	var configData dto.ConfigData
	if err := godotenv.Load(".env"); err != nil {
		return configData, err
	}

	if port := os.Getenv("PORT"); port != "" {
		configData.AppConfig.Port = port
	}

	dbHost := os.Getenv("DB_PORT")
	dbPort := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbMaxIdle := os.Getenv("MAX_IDLE")
	dbMaxConn := os.Getenv("MAX_CONN")
	dbMaxLifeTime := os.Getenv("MAX_LIFE_TIME")
	dbLogMode := os.Getenv("LOG_MODE")
	saltInt := os.Getenv("SALT")
	secretToken := os.Getenv("SECRET_TOKEN")
	tokenExpired := os.Getenv("TOKEN_EXPIRED")
	Roles := os.Getenv("ROLES")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbMaxIdle == "" || dbMaxConn == "" || dbMaxLifeTime == "" || dbLogMode == "" || saltInt == "" || secretToken == "" || tokenExpired == "" || Roles == "" {
		return configData, errors.New("DB Config not set")
	}

	var err error
	configData.DbConfig.MaxConn, err = strconv.Atoi(dbMaxConn)
	if err != nil {
		return configData, err
	}

	configData.DbConfig.MaxIdle, err = strconv.Atoi(dbMaxIdle)
	if err != nil {
		return configData, err
	}
	configData.DbConfig.Host = dbHost
	configData.DbConfig.DbPort = dbPort
	configData.DbConfig.User = dbUser
	configData.DbConfig.Pass = dbPassword
	configData.DbConfig.Database = dbName
	configData.DbConfig.MaxLifeTime = dbMaxLifeTime
	configData.DbConfig.LogMode, err = strconv.Atoi(dbLogMode)
	configData.DbConfig.Salt, err = strconv.Atoi(saltInt)
	configData.DbConfig.SecretToken = secretToken
	configData.DbConfig.TokenExpire, err = strconv.Atoi(tokenExpired)
	configData.DbConfig.Roles = Roles
	if err != nil {
		return configData, err
	}

	return configData, nil
}

func RunService() {
	// Create log file
	if err := createLogFile(); err != nil {
		log.Fatal().Msg("Error creating log file: " + err.Error())
	}

	//add log
	zerolog.TimeFieldFormat = "02-01-2006 15:04:05"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	//load env config
	configData, err := initEnv()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	conn, err := config.ConnectToDB(configData, log.Logger)
	if err != nil {
		log.Error().Msg("service.ConnectToDB.Err: " + err.Error())
		return
	}

	duration, err := time.ParseDuration(configData.DbConfig.MaxLifeTime)
	if err != nil {
		log.Error().Msg("service.Duration.Err: " + err.Error())
		return
	}

	utils.InitConfigData(configData)
	middleware.InitConfigData(configData)

	conn.SetConnMaxLifetime(duration)
	conn.SetMaxIdleConns(configData.DbConfig.MaxIdle)
	conn.SetMaxOpenConns(configData.DbConfig.MaxConn)

	defer func() {
		errClose := conn.Close()
		if errClose != nil {
			log.Error().Msg(errClose.Error())
		}
	}()

	// set up time zone
	time.Local = time.FixedZone("Asia/Jakarta", 7*60*60)
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins:    []string{"*"},
		AllowMethods:    []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin", "Content-type", "Authorization",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           120 * time.Second,
	}))

	log.Logger = log.With().Caller().Logger()

	r.Use(logger.SetLogger(
		logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(os.Stdout).With().Logger()
		}),
	))

	//gin recovery for handle panic
	r.Use(gin.Recovery())
	r.Use(RequestLog())
	initializeDomainModule(r, conn)

	version := "0.0.1"
	log.Info().Msg(fmt.Sprintf("Service running versions %s", version))
	addr := flag.String("Port: ", ":"+configData.AppConfig.Port, "Address to listen and serve")
	err = r.Run(*addr)
	if err != nil {
		log.Error().Msg(err.Error())
	}

}

func initializeDomainModule(r *gin.Engine, db *sql.DB) {
	apiGroup := r.Group("/api")
	v1Group := apiGroup.Group("/v1")

	router.InitRouter(v1Group, db)
}

func createLogFile() error {
	dir := "./.log/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	today := time.Now().Format("2006-01-02")

	file, err := os.OpenFile(dir+today+".json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, file)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	return nil
}

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		log.Info().
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Str("user_agent", c.Request.UserAgent()).
			Int("body_size", c.Writer.Size()).
			Msg("Request")
	}
}
