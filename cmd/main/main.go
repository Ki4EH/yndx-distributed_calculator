package main

import (
	"fmt"
	"github.com/Ki4EH/yndx-distributed_calculator/internal/handlers"
	"github.com/Ki4EH/yndx-distributed_calculator/internal/logger"
	"net/http"
)

func main() {
	http.Handle("/expression", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handlers.AddExpressionHandler(writer, request)
	}))
	log, _ := logger.NewFileLogger()
	log.Info("[SERVER] Starting server on localhost:8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Error(fmt.Sprintf("Error starting server: %v", err))
	}
	defer log.Close()
}
