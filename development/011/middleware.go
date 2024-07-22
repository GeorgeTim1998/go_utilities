package main

import (
	"log"
	"net/http"
	"time"
)

// Middleware для логирования запросов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Запрос: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Ответ: %s %s завершен за %v", r.Method, r.URL.Path, time.Since(start))
	})
}
