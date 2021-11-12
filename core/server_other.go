//go:build !windows
// +build !windows

package core

import (
	"crypto/tls"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	s.TLSConfig = &tls.Config{NextProtos: []string{"h2"}}
	return s
}
