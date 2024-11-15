package service

import (
	"context"
	"errors"
	"go-project/pkg/util"
)

// 接口逻辑

type Service struct {
	ctx context.Context
}

func NewService() *Service {
	return &Service{
		ctx: context.Background(),
	}
}

func (s *Service) Hello(msg string) (string, util.RespError) {
	switch msg {
	case "a":
		return "a", nil
	case "b":
		return "b", util.NewRespError(errors.New("test error"), true)
	}
	return "hello", nil
}
