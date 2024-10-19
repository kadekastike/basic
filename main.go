package main

import (
	"encoding/json"
	"learnApi/controllers"
	"learnApi/middleware"
	"learnApi/models"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	models.ConnectToDatabase()

	routerMiddleware := middleware.RequestResponseLogger(router)

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "hello world",
		})
	})

	router.HandleFunc("GET /api/posts", controllers.FindPosts)
	router.HandleFunc("POST /api/post", controllers.StorePost)
	router.HandleFunc("GET /api/post/{id}", controllers.FindPostById)
	router.HandleFunc("PUT /api/post/{id}", controllers.UpdatePost)
	router.HandleFunc("DELETE /api/post/{id}", controllers.DeletePost)

	http.ListenAndServe(":8000", routerMiddleware)
}
