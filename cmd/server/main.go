package main

import (
	"fmt"
	"go-microservice-app/pkg/api"
	"go-microservice-app/pkg/db"
	"log"
	"net/http"
	"os"
)

func main()  {
  pgdb, err := db.StartDB()

  if err != nil {
    log.Printf("error starting the database %v\n", err)
    os.Exit(4)
  }
  log.Printf("Sucessfully sync with DATABASE...\n")


  router := api.StartAPI(pgdb)

  port := os.Getenv("PORT")

  log.Printf("the server is listening on the port: http://127.0.0.1:%s...\n", port)
  err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)

  if err != nil {
    log.Printf("error from router %v\n", err)
    os.Exit(4)
  }

}
