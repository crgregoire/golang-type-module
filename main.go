package main

import (
	"flag"
	"os"

	"github.com/subosito/gotenv"
	"github.com/tespo/satya/v2/migrations"
	"github.com/tespo/satya/v2/seeders"
)

func init() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "local"
	}
	gotenv.Load("./config/" + env + ".env")
}

func main() {
	t := flag.String("type", "migration", "Used to run specific functions")
	flag.Parse()
	if *t == "migrate" {
		migrations.Migrate()
	}
	if *t == "seed" {
		seeders.SeedDatabase()
	}
	if *t == "nuke" {
		migrations.Nuke()
	}
	if *t == "build" {
		migrations.Migrate()
		seeders.SeedDatabase()
	}
}
