/*
Copyright © 2022 ANDREY VALITOV <andreivalitov2001@gmail.com>
*/

package commands

import (
	"log"
	"migrator/internal/commons"
	"migrator/internal/configs"
	"strconv"

	"github.com/spf13/cobra"
)

func NewDownCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "down [ver]",
		Short: "Down migration operation that makes the database version the initially or corresponding to a given version",
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
				ver *= -1
				err = migrator.Steps(ver)
				commons.HandleMigrationErr(err, "Down complete!")
				return
			}

			if len(args) != 0 {
				log.Fatal("down without -v flag does not accept any parameters")
			}
			err = migrator.Down()
			commons.HandleMigrationErr(err, "Down complete!")
		},
	}

	command.PersistentFlags().BoolP("with-version", "v", false, "migrate with specified version")
	return command
}
