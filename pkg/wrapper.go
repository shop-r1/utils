package pkg

import (
	"context"
	"github.com/micro/go-micro/server"
)

// logWrapper is a handler wrapper
func Authorization(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		//log.Printf("[wrapper] server request: %v", req.Endpoint())
		//md, _ := metadata.FromContext(ctx)
		err := fn(ctx, req, rsp)
		return err
	}
}
