package rpcxx

import "errors"

var (
	ErrRPCServerDisabled = errors.New("server is disabled")
	ErrInvalidToken      = errors.New("invalid token")
)
