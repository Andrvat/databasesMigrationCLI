/*
Copyright © 2022 ANDREY VALITOV <andreivalitov2001@gmail.com>
*/

package cmd

import (
	"log"
	"migrator/migrating"
	"strconv"

	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down [ver]",
	Short: "Down migration operation that makes the database version the initially or corresponding to a given version",
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
			ver *= -1
			err = migrator.Steps(ver)
		} else {
			if len(args) != 0 {
				log.Fatal("down without -v flag does not accept any parameters")
			}
			err = migrator.Down()
		}
		migrating.Finally(err)
		log.Println("Down complete!")
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	downCmd.PersistentFlags().BoolP("with-version", "v", false, "migrate with specified version")
}
