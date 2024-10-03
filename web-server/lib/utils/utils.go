package utils

import "errors"

var (
	ErrServerRunning  = errors.New("server is already running")
	ErrServerStopping = errors.New("server is not running")
)
