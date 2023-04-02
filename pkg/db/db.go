package db

import (
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
	if os.Getenv("ENV") == "PROD" {
		opts, err = pg.ParseURL(os.Getenv("DATABASE_URL"))
		if err != nil {
			return nil, err
		}
	} else {
		opts = &pg.Options{
			//default port
			//depends on the db service from docker compose
			Addr:     "db:5432",
			User:     "postgres",
			Password: "postgres",
			Database: "postgres",

		}
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