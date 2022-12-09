// Code generated by microgen 0.10.0. DO NOT EDIT.

package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	service "github.com/valerylobachev/microgen/examples/usersvc/pkg/usersvc"
)

// ErrorLoggingMiddleware writes to logger any error, if it is not nil.
func ErrorLoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.UserService) service.UserService {
		return &errorLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type errorLoggingMiddleware struct {
	logger log.Logger
	next   service.UserService
}

func (M errorLoggingMiddleware) CreateUser(ctx context.Context, user service.User) (id string, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "CreateUser", "message", err)
		}
	}()
	return M.next.CreateUser(ctx, user)
}

func (M errorLoggingMiddleware) UpdateUser(ctx context.Context, user service.User) (err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "UpdateUser", "message", err)
		}
	}()
	return M.next.UpdateUser(ctx, user)
}

func (M errorLoggingMiddleware) GetUser(ctx context.Context, id string) (user service.User, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GetUser", "message", err)
		}
	}()
	return M.next.GetUser(ctx, id)
}

func (M errorLoggingMiddleware) FindUsers(ctx context.Context) (results map[string]service.User, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "FindUsers", "message", err)
		}
	}()
	return M.next.FindUsers(ctx)
}

func (M errorLoggingMiddleware) CreateComment(ctx context.Context, comment service.Comment) (id string, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "CreateComment", "message", err)
		}
	}()
	return M.next.CreateComment(ctx, comment)
}

func (M errorLoggingMiddleware) GetComment(ctx context.Context, id string) (comment service.Comment, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GetComment", "message", err)
		}
	}()
	return M.next.GetComment(ctx, id)
}

func (M errorLoggingMiddleware) GetUserComments(ctx context.Context, userId string) (list []service.Comment, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GetUserComments", "message", err)
		}
	}()
	return M.next.GetUserComments(ctx, userId)
}
