package controllers

import (
	"api-gin2/configs"
	"api-gin2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users []models.Users

// fazendo requisição por where email
func FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	configs.DB.Where("email = ?", id).Find(&users)

	ctx.JSON(http.StatusOK, &users)
}

// fazendo requisição pegando todos da tabela
func FindAll(ctx *gin.Context) {
	configs.DB.Find(&users)

	ctx.JSON(http.StatusOK, &users)
}

// create user
func CreateTweet(ctx *gin.Context) {

	var user models.Users

	if err := ctx.BindJSON(&user); err != nil {
		return
	}

	result := configs.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, user)

}

func UpdateTweet(ctx *gin.Context) {

	var user models.Users
	id := ctx.Param("id")

	if err := ctx.BindJSON(&user); err != nil {
		return
	}
	// configs.DB.Where("email = ?", id).Find(&user)
	// result := configs.DB.Model(&user).Updates(&user)

	result := configs.DB.Model(&user).Where("email = ?", id).Updates(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
		return
	}

	ctx.JSON(http.StatusOK, user)

}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	result := configs.DB.Delete(&users, "email = ?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Successfully deleted user"})

}
