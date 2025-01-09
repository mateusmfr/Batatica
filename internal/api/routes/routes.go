package routes

import (
	"batatica/internal/api/handlers"
	"net/http"
)

func InitializeRoutes() {
	http.HandleFunc("/api/test", handlers.TestHandler)
}
