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
	res, err := svc.Hello()
	if err != nil {
		logrus.Errorln(err)
		util.Respf().Fail(xresponse.HttpInternalServerError).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithMsg(res).Response(util.GinRespf(c))
}
