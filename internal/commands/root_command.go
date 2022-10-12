/*
Copyright © 2022 ANDREY VALITOV <andreivalitov2001@gmail.com>
*/

package commands

import (
	"github.com/spf13/cobra"
	"log"
)

type MigrationCli struct {
	rootCommand *cobra.Command
}

func NewMigrationCli() *MigrationCli {
	rootCommand := &cobra.Command{
		Use:   "migrator",
		Short: "CLI application that provides databases migration operations",
		Long: `CLI application that provides databases migration operations {up [ver], down [ver]}
Before run this CLI app you need to export the following environment variables:
DRIVER_NAME							Example: postgres
DATABASE_USER							Example: username
DATABASE_PASSWORD						Example: your123password
DATABASE_URL		Format: url:port			Example: localhost:5432
DATABASE_NAME							Example: awesome_db
SOURCE_URL		Format: path/to/migrations/folder 	Example: /home/golang/project/migrations`,
	}

	rootCommand.AddCommand(NewDownCommand())
	rootCommand.AddCommand(NewUpCommand())
	return &MigrationCli{rootCommand}
}

func (m *MigrationCli) Execute() {
	err := m.rootCommand.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
