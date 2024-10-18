package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"learnApi/models"
	"learnApi/utils"
	"learnApi/validation"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func FindPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	models.DB.Find(&posts)

	response := map[string]interface{}{
		"success": true,
		"message": "List post success",
		"posts":   posts,
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this field is required"
	case "min":
		return "minimum length is 3"
	case "max":
		return "maximal length is 100"
	}

	fmt.Println(fe.Tag())
	return "Unknown Error"

}

func StorePost(w http.ResponseWriter, r *http.Request) {
	var input validation.ValidatePostInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]validation.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = validation.ErrorMsg{Field: fe.Field(), Message: GetErrorMsg(fe)}
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": out})
		}
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}

	models.DB.Create(&post)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "post created!",
		"data":    post,
	})

}

func FindPostById(w http.ResponseWriter, r *http.Request) {
	post,err := utils.FindDataById(w, r)
	if err != nil {
		return
	}
	w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{} {
			"success": true,
			"message": "Data found",
			"data":    post,
		})
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	post, err := utils.FindDataById(w,r)
	if err != nil {
		return
	}
	var input validation.ValidatePostInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]validation.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = validation.ErrorMsg{Field: fe.Field(), Message: GetErrorMsg(fe)}
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": out})
		}
		return
	}

	models.DB.Model(&post).Updates(input)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Post Updated Successfully",
		"data":    post,
	})
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	post, err := utils.FindDataById(w,r)
	if err != nil {
		return
	}
	models.DB.Delete(&post)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Post Deleted Successfully",
	})
}
