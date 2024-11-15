package configpost

import (
	"context"
	"go-project/pkg/config"
	"go-project/server/controller"
)

func PostInit(ctx context.Context) {
	// 启动 HTTP 服务
	controller.RunServer(ctx, config.Conf.ServerBind)
}
