package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func GetValidate[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body T
		if err := c.ShouldBindQuery(&body); err != nil {
			println(reflect.TypeOf(err.Error()))
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}
		c.Next()
	}
}

func PostJsonValidate[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body T
		if err := c.ShouldBindJSON(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}
		c.Next()
	}
}

func PostFormValidate[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body T
		if err := c.ShouldBind(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}
		c.Next()
	}
}

func UriValidate[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body T
		fmt.Println(&body)
		if err := c.ShouldBindUri(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}
		c.Next()
	}
}
