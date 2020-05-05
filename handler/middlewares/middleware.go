package middlewares

import (
	"fmt"
	"net/http"
)

func Middleware1(next http.Handler) http.Handler {
	fmt.Println("[SET] Middleware1")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] Middleware1")
		next.ServeHTTP(w, r)
		fmt.Println("[END] Middleware1")
	})
}

func Middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] Middleware2")
		next.ServeHTTP(w, r)
		fmt.Println("[END] Middleware2")
	})
}

func Middleware3(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] Middleware3")
		next.ServeHTTP(w, r)
		fmt.Println("[END] Middleware3")
	})
}
