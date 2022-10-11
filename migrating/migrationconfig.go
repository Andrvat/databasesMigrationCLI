package migrating

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"os"
)

type MigratorConfig struct {
	DriverName       string `validate:"required"`
	DatabaseUser     string `validate:"required"`
	DatabasePassword string `validate:"required"`
	DatabaseUrl      string `validate:"required"`
	DatabaseName     string `validate:"required"`
	SourceUrl        string `validate:"required"`
}
// bluh
var validate *validator.Validate

func newMigratorConfig() (*MigratorConfig, error) {
	config := &MigratorConfig{
		DriverName:       os.Getenv("DRIVER_NAME"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseUrl:      os.Getenv("DATABASE_URL"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		SourceUrl:        os.Getenv("SOURCE_URL"),
	}

	validate = validator.New()
	err := validate.Struct(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func ConfigureMigrator() (*migrate.Migrate, *sql.DB, error) {
	config, err := newMigratorConfig()
	if err != nil {
		return nil, nil, err
	}

	db, err := sql.Open(config.DriverName,
		fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
			config.DriverName, config.DatabaseUser, config.DatabasePassword, config.DatabaseUrl, config.DatabaseName))
	if err != nil {
		return nil, nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, db, err
	}
	migrator, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", config.SourceUrl), "", driver)
	if err != nil {
		return nil, db, err
	}

	return migrator, db, nil
}
