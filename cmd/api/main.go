package main

import (
	"fmt"
	"log"
	"net/http"

	"batatica/internal/api/routes"
)

func main() {
	routes.InitializeRoutes()

	port := ":8080"
	fmt.Println("Servidor rodando na porta", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
