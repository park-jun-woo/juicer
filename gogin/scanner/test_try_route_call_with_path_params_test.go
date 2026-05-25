//ff:func feature=scan type=extract control=sequence
//ff:what TestTryRouteCall_WithPathParams 테스트
package scanner

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestTryRouteCall_WithPathParams(t *testing.T) {
	src := `package main

import "github.com/gin-gonic/gin"

func setup() {
	r := gin.Default()
	r.GET("/users/:id", handler)
}

func handler(c *gin.Context) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "main.go", src, 0)
	eps := scanFile(file, "main.go", fset)
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if eps[0].Request == nil || len(eps[0].Request.PathParams) != 1 {
		t.Error("expected 1 path param")
	}
}
