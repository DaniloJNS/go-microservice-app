package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
)

func StartDB() (*pg.DB, error) {
	var (
		opts *pg.Options
		err  error
	)

	log.Printf("Starting connection with DATABASE...\n")

	opts = &pg.Options{
		Addr:     fmt.Sprint(getEnvWithDefault("POSTGRES_HOST", "db"), ":", getEnvWithDefault("POSTGRES_PORT", "5432")),
		User:     getEnvWithDefault("POSTGRES_USER", "postgres"),
		Password: getEnvWithDefault("POSTGRES_PASSWORD", "postgres"),
		Database: getEnvWithDefault("POSTGRES_DATABASE", "postgres"),
	}

	// Start connection in DB
	db := pg.Connect(opts)

	// run migrations
	collection := migrations.NewCollection()
	err = collection.DiscoverSQLMigrations("migrations")

	if err != nil {
		return nil, err
	}

	//start the migrations
	_, _, err = collection.Run(db, "init")

	if err != nil {
		return nil, err
	}

	oldVersion, newVersion, err := collection.Run(db, "up")

	if err != nil {
		return nil, err
	}

	if newVersion != oldVersion {
		log.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		log.Printf("version is %d\n", oldVersion)
	}

	// return the db connection
	return db, err
}

func getEnvWithDefault(envName string, defaultValue string) string {
	envValue, envDefined := os.LookupEnv(envName)
	if envDefined {
		return envValue
	}
	return defaultValue
}
