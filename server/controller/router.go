package controller

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RunServer(ctx context.Context, addr string) {
	r := gin.Default()
	router(r)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			if err != nil {
				logrus.Fatalf("server run error: %s", err)
			}
		}
	}()
	<-ctx.Done()
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("server shutdown error: %s", err)
	}
	logrus.Warn("server has been shutdown")
}

func router(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/hello", Hello)
	}
}
