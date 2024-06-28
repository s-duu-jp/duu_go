package impl

import "api/handlers"

type handlersImpl struct {
}

var _ handlers.AuthenticationHandlers = &handlersImpl{}

func NewHandlers() *handlersImpl {
	return &handlersImpl{}
}
