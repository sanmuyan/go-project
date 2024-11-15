package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"go-project/pkg/util"
	"go-project/server/service"
)

// API列表

var svc = service.NewService()

func Hello(c *gin.Context) {
	msg := c.Query("msg")
	res, err := svc.Hello(msg)
	if err != nil {
		logrus.Errorf("error: %v", err)
		util.Respf().Fail(xresponse.HttpInternalServerError).WithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithMsg(res).Response(util.GinRespf(c))
}
