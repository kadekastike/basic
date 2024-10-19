package utils

import (
	"encoding/json"
	"learnApi/models"
	"net/http"
)

func FindDataById(w http.ResponseWriter, r *http.Request) (*models.Post, error) {
	var post models.Post
	
	id := r.PathValue("id")

	if err := models.DB.Where("id = ?", id).First(&post).Error; err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Data Not found",
		})
		return nil, err
	}

	return &post, nil

}
