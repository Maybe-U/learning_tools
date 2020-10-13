package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
)

//中间件实现方式 ，实现 endpoint.Middleware
func NewGolangRateAllowMiddleware(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return "", errors.New("limit req  Allow")
			}
			return next(ctx, request)
		}
	}
}

//中间件实现方式最简单的方式，对endpoint.Endpoint进行封装
func MyMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return next(ctx, request)
	}
}
