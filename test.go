package main

import (
	"github.com/gin-gonic/gin"
	"go.opencensus.io/plugin/ochttp"
	"net/http"
)

type Bind struct {
	Name string `json:"name" binding:"required"`
}

type BindParam struct {
	Name string `uri:"name" binding:"required"`
}

type BindQuery struct {
	Name string `form:"name" binding:"required"`
}

func main() {
	r := gin.New()
	r.POST("/", func(c *gin.Context) {
		req := &Bind{}
		err := c.Bind(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.Status(http.StatusOK)
	})
	r.POST("/:name", func(c *gin.Context) {
		req := &BindParam{}
		err := c.ShouldBindUri(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		c.Status(http.StatusOK)
	})
	r.GET("/", func(c *gin.Context) {
		req := &BindQuery{}
		err := c.BindQuery(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		c.Status(http.StatusOK)
	})
	http.ListenAndServe("localhost:9999", &ochttp.Handler{Handler: r})
}
