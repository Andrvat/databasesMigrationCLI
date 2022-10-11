/*
Copyright © 2022 ANDREY VALITOV <andreivalitov2001@gmail.com>
*/

package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
