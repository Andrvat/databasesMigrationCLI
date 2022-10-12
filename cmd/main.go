/*
Copyright © 2022 ANDREY VALITOV <andreivalitov2001@gmail.com>
*/
package main

import (
	"migrator/internal/commands"
)

func main() {
	cli := commands.NewMigrationCli()
	cli.Execute()
}
