//ff:func feature=scan type=extract control=sequence
//ff:what TestScanFile_WithParamRouter 테스트
package gogin

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestScanFile_WithParamRouter(t *testing.T) {
	src := `package main

import "github.com/gin-gonic/gin"

func setup(r *gin.Engine) {
	r.GET("/users", handler)
}

func handler(c *gin.Context) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "main.go", src, 0)
	eps, _ := scanFile(file, "main.go", fset)
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
}
