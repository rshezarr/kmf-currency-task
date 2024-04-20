package database

import (
	"context"
	"currency/internal/config"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func InitializeDB(dbCfg config.DBConfig) *sqlx.DB {
	// Подключение к БД
	db, err := sqlx.Open(dbCfg.DriverName, dbCfg.GetConnectionString())
	if err != nil {
		zap.S().Fatalf("Error connecting to database:%v", err)
	}

	// Проверка подключения
	err = db.PingContext(context.Background())
	if err != nil {
		zap.S().Fatalf("Error pinging database: %v", err)
	}

	zap.S().Info("connected to the database")
	return db
}
