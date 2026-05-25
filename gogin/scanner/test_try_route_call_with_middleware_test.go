//ff:func feature=scan type=extract control=sequence
//ff:what TestTryRouteCall_WithMiddleware 테스트
package scanner

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestTryRouteCall_WithMiddleware(t *testing.T) {
	src := `package main

import "github.com/gin-gonic/gin"

func setup() {
	r := gin.Default()
	r.Use(mw)
	r.GET("/test", handler)
}

func mw(c *gin.Context) {}
func handler(c *gin.Context) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "main.go", src, 0)
	eps := scanFile(file, "main.go", fset)
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if len(eps[0].Middleware) != 1 {
		t.Errorf("expected 1 middleware, got %d", len(eps[0].Middleware))
	}
}
