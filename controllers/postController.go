package controllers

import (
	"errors"
	"learnApi/models"
	"learnApi/validation"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func FindPosts(ctx *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "List post success",
		"posts":   posts,
	})
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this field is required"
	}
	return "Unknown Error"

}

func StorePost(ctx *gin.Context) {
	var input validation.ValidatePostInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]validation.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = validation.ErrorMsg{Field: fe.Field(), Message: GetErrorMsg(fe)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}

	models.DB.Create(&post)

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "Post Created",
		"data":    post,
	})

}

func FindPostById(ctx *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found"})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "Data found",
		"data":    post,
	})
}

func UpdatePost(ctx *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input validation.ValidatePostInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]validation.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = validation.ErrorMsg{Field: fe.Field(), Message: GetErrorMsg(fe)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	models.DB.Model(&post).Updates(input)

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "Post Updated Successfully",
		"data":    post,
	})
}

func DeletePost(ctx *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//delete post
	models.DB.Delete(&post)

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "Post Deleted Successfully",
	})
}