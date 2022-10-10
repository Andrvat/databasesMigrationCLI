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
	Run: func(cmd *cobra.Command, args []string) {
		withVersion, _ := cmd.Flags().GetBool("with-version")
		if withVersion {
			ver, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
			downWithVersion(ver)
		} else {
			down()
		}
		log.Println("Down complete!")
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	downCmd.PersistentFlags().BoolP("with-version", "v", false, "migrate with specified version")
}

func down() {
	migrator, db, err := migrating.ConfigureMigrator()
	defer migrating.CloseConn(db)
	if err != nil {
		log.Fatal(err)
	}

	err = migrator.Down()
	migrating.Finally(err)
}

func downWithVersion(ver int) {
	migrator, db, err := migrating.ConfigureMigrator()
	defer migrating.CloseConn(db)
	if err != nil {
		log.Fatal(err)
	}

	ver *= -1
	err = migrator.Steps(ver)
	migrating.Finally(err)
}
