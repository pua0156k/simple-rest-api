package controller

import (
	"net/http"
	"simpleRestAPI/model"
	"simpleRestAPI/service"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	var comment model.Comment
	err := c.ShouldBindJSON(&comment)
	comments, err := service.Insert(comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	} else {
		c.JSON(200, comments)
	}

}

func GetComment(c *gin.Context) {
	result, err := service.FindById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	} else {
		c.JSON(200, result)
	}

}

func UpdateComment(c *gin.Context) {
	var comment model.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

	_, err = service.UpdateById(c.Param("id"), comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	} else {
		c.JSON(200, comment)
	}

}

func DeleteComment(c *gin.Context) {
	_, err := service.DeleteById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	} else {
		c.JSON(200, "")
	}

}
