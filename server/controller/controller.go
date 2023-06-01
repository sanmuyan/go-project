package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-project/pkg/response"
	"go-project/server/service"
)

// API列表

var svc = service.NewService()

var respf = func() *response.Response {
	return response.NewResponse()
}

func Hello(c *gin.Context) {
	res, err := svc.Hello()
	if err != nil {
		logrus.Errorln(err)
		respf().Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	respf().Ok().WithMsg(res).SetGin(c)
}
