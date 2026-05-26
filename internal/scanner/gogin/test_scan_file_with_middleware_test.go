//ff:func feature=scan type=extract control=sequence
//ff:what TestScanFile_WithMiddleware 테스트
package gogin

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestScanFile_WithMiddleware(t *testing.T) {
	src := `package main

import "github.com/gin-gonic/gin"

func setup() {
	r := gin.Default()
	r.Use(authMiddleware)
	r.GET("/users", handler)
}

func authMiddleware(c *gin.Context) {}
func handler(c *gin.Context) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "main.go", src, 0)
	eps, _ := scanFile(file, "main.go", fset)
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if len(eps[0].Middleware) != 1 {
		t.Errorf("expected 1 middleware, got %d", len(eps[0].Middleware))
	}
}
