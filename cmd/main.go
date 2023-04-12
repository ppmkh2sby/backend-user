package main

import (
	"fmt"
	"github.com/ppmkh2sby/backend-user/config"
	_db "github.com/ppmkh2sby/backend-user/db"
	"github.com/sirupsen/logrus"
	"runtime"
)

const (
	postgresDB = "postgres"
)

func main() {
	//custom format log
	customFormatter := &logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
		CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" %s:%d", frame.File, frame.Line)
		},
	}
	logrus.SetFormatter(customFormatter)
	logrus.SetReportCaller(true)

	//get config
	configLoader := config.NewYamlConfigLoader("backend-user.yaml")
	config, err := configLoader.GetServiceConfig()
	if err != nil {
		logrus.Fatalf("Unable to load configuration: %w", err)
	}

	// Initializa database connection
	db, err := _db.InitPostgresDB(postgresDB, config.SourceData.PostgresDBServer, config.SourceData.PostgresDBName,
		config.SourceData.PostgresDBUsername, config.SourceData.PostgresDBPassword, config.SourceData.PostgresDBPort,
		config.SourceData.PostgresDBTimeout)
	if err != nil {
		logrus.Fatalf("Unable to create db instance: %v", err)
	}
	logrus.Infof("Success create postgreDB instance")
}
