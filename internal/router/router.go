package router

import (
	"fmt"
	"labForBosz/internal/handler"
	"labForBosz/pkg/model/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(hdl handler.Handle) *gin.Engine {
	r := gin.Default()

	r.POST("/customer", func(c *gin.Context) {
		var req model.CreateCustomerRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		id, err := hdl.CreateCustomer(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	r.GET("/customer/:id", func(c *gin.Context) {
		var req model.GetCustomerReqeust
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Printf("\n\nreq: %+v\n\n: ", req)

		resp, err := hdl.GetCustomer(c.Request.Context(), req.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	r.GET("/customer/:id/name", func(c *gin.Context) {
		var req model.GetCustomerReqeust
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		resp, err := hdl.GetCustomerName(c, req.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	return r
}
