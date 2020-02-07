package main

import (
	"flag"
	"os"
	"testing"

	"github.com/tespo/satya/v2/migrations"
	"github.com/tespo/satya/v2/seeders"
)

func TestMain(m *testing.M) {
	local := flag.String("local", "false", "Determines best way to run tests")
	flag.Parse()
	os.Setenv("DB_NAME", "tespo_docker")
	os.Setenv("DB_USER", "root")
	if *local != "true" {
		migrations.Migrate()
		seeders.SeedDatabase()
	}
	result := m.Run()
	os.Exit(result)
}
