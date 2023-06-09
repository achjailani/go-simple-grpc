package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloHandler struct {
	*Handler
}

type HelloRequest struct {
	text string
}

func NewHelloHandler(h *Handler) *HelloHandler {
	return &HelloHandler{h}
}

// SayHello is a method to handle
func (hdl *HelloHandler) SayHello(c *gin.Context) {
	var body HelloRequest
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	rsp, err := hdl.Dependency.GRPCClient.SayHello(c, body.text)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": rsp.Message,
	})
	return
}

// Ping is a method to handle
func (hdl *HelloHandler) Ping(c *gin.Context) {
	rsp, err := hdl.Dependency.GRPCClient.Ping(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": rsp,
	})
	return
}
