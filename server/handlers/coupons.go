package handlers

import (
	s "OnlineStoreBackend/server"
)

type HandlersCoupons struct {
	server *s.Server
}

func NewHandlersCoupons(server *s.Server) *HandlersCoupons {
	return &HandlersCoupons{server: server}
}
