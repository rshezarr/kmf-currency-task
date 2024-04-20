package database

import (
	"context"
	"currency/internal/config"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"go.uber.org/zap"
)

func InitializeDB(dbCfg config.DBConfig) *sql.DB {
	// Подключение к БД
	zap.S().Info(dbCfg.GetConnectionString())
	db, err := sql.Open(dbCfg.DriverName, dbCfg.GetConnectionString())
	if err != nil {
		zap.S().Fatalf("Error connecting to database:%v", err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.PingContext(context.Background())
	if err != nil {
		zap.S().Fatalf("Error pinging database: %v", err)
	}

	zap.S().Info("connected to the database")
	return db
}
