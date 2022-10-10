/*
Copyright © 2022 ANDREY VALITOV <andreivalitov2001@gmail.com>
*/

package cmd

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"log"
	"migrator/migrating"
	"strconv"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up [ver]",
	Short: "Up migration operation that makes the database version the latest or corresponding to a given version",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		migrator, db, err := migrating.ConfigureMigrator()
		defer migrating.CloseConn(db)
		if err != nil {
			log.Fatal(err)
		}

		withVersion, _ := cmd.Flags().GetBool("with-version")
		if withVersion {
			ver, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
			err = migrator.Steps(ver)
		} else {
			if len(args) != 0 {
				log.Fatal("up without -v flag does not accept any parameters")
			}
			err = migrator.Up()
		}
		migrating.Finally(err)
		log.Println("Up complete!")
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	upCmd.PersistentFlags().BoolP("with-version", "v", false, "migrate with specified version")
}
