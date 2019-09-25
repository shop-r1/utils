package wrappers

import (
	"context"
	"github.com/micro/go-micro/server"
	log "github.com/sirupsen/logrus"
)

func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[wrapper] server request: %v \n body: %v\n", req.Endpoint(), req.Body())
		err := fn(ctx, req, rsp)
		return err
	}
}
