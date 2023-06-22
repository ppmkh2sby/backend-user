package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ppmkh2sby/backend-library/helpers/image"
	"github.com/ppmkh2sby/backend-user/app/route"
	"github.com/ppmkh2sby/backend-user/config"
	"github.com/ppmkh2sby/backend-user/db"
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
	logrus.Printf("Developed by %v\n%v", image.Developer, image.Ppmkh2sbyASCII)
	logrus.Println("Start Service...")

	//get config
	configLoader := config.NewYamlConfigLoader("backend_user.yaml")
	config, err := configLoader.GetServiceConfig()
	if err != nil {
		logrus.Fatalf("Unable to load configuration: %w", err)
	}

	// Initialize database connection
	db, err := db.InitPostgresDB("postgres", config.SourceData.PostgresDBServer, config.SourceData.PostgresDBName,
		config.SourceData.PostgresDBUsername, config.SourceData.PostgresDBPassword, config.SourceData.PostgresDBPort,
		config.SourceData.PostgresDBTimeout)
	if err != nil {
		logrus.Fatalf("Unable to create db instance: %v", err)
	}
	logrus.Infof("Success create postgreDB instance")

	// Routing app with fiber
	app := fiber.New()
	route.Init(app, db)
	logrus.Infof("Success routing app...")

	// Running app
	err = app.Listen(config.ServiceData.Address)
	if err != nil {
		logrus.Fatalf("Unable to run http server: %v", err)
	}
	logrus.Infof("Service be running at: %v", config.ServiceData.Address)
}
