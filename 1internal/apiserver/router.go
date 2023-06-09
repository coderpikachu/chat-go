// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"chat-go/1internal/pkg/middleware/auth"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}
func printHello(c *gin.Context) {
	core.WriteResponse(c, nil, "hello")
}

func installController(g *gin.Engine) *gin.Engine {
	// Middlewares.
	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)
	g.POST("/hello", printHello)
	return g
}
