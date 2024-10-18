package utils

import (
	"encoding/json"
	"learnApi/models"
	"net/http"
	"strings"
)

func FindDataById(w http.ResponseWriter, r *http.Request)(*models.Post, error) {
	var post models.Post
	path := r.URL.Path
	parts := strings.Split(path, "/")
		if len(parts) < 4 || parts[3] == "" {
			http.Error(w, `{"error": "Invalid ID"}`, http.StatusBadRequest)
			return nil, nil
		}
	
	id := parts[3] 

	if err := models.DB.Where("id = ?", id).First(&post).Error; err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{} {
			"success": false,
			"message": "Data Not found",
		})
		return nil, err
	}

	return &post, nil

}