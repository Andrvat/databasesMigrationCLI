package configs

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"os"
)

type MigratorConfig struct {
	DataSourceName string `validate:"required,contains=sslmode=disable"`
	SourceUrl      string `validate:"required"`
}

var validate *validator.Validate

func newMigratorConfig() (*MigratorConfig, error) {
	config := &MigratorConfig{
		DataSourceName: os.Getenv("POSTGRES_DSN"),
		SourceUrl:      getEnvOrDefault("SOURCE_URL", "./"),
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

	db, err := sql.Open("postgres", config.DataSourceName)
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

func getEnvOrDefault(envKey string, defaultValue string) string {
	if value, exist := os.LookupEnv(envKey); exist {
		return value
	}
	return defaultValue
}
