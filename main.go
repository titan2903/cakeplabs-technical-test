package main

import (
	"cakeplabs-technical-test/pkg/utils"
	"cakeplabs-technical-test/routes"
	"fmt"
)

func main() {
	e := routes.NewRoutes()
	port := utils.GoDotEnvVariable("PORT")
	if port == "" {
		port = "8080"
	}

	sPort := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(sPort))
	fmt.Printf("Successfully started on port %s\n", port)
}
