/*
Copyright © 2022 ANDREY VALITOV <andreivalitov2001@gmail.com>
*/

package commands

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"log"
	"migrator/internal/commons"
	"migrator/internal/configs"
	"strconv"
)

func NewUpCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "up [ver]",
		Short: "Up migration operation that makes the database version the latest or corresponding to a given version",
		Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			migrator, db, err := configs.ConfigureMigrator()
			if err != nil {
				log.Fatal(err)
			}
			defer commons.CloseConn(db)

			withVersion, err := cmd.Flags().GetBool("with-version")
			if err != nil {
				log.Fatal(err)
			}

			if withVersion {
				ver, err := strconv.Atoi(args[0])
				if err != nil {
					log.Fatal(err)
				}
				err = migrator.Steps(ver)
				commons.HandleMigrationErr(err, "Up complete!")
				return
			}

			if len(args) != 0 {
				log.Fatal("up without -v flag does not accept any parameters")
			}
			err = migrator.Up()
			commons.HandleMigrationErr(err, "Up complete!")
		},
	}

	command.PersistentFlags().BoolP("with-version", "v", false, "migrate with specified version")
	return command
}
