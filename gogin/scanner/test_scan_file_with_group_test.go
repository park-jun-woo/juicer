//ff:func feature=scan type=extract control=sequence
//ff:what TestScanFile_WithGroup 테스트
package scanner

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestScanFile_WithGroup(t *testing.T) {
	src := `package main

import "github.com/gin-gonic/gin"

func setup() {
	r := gin.Default()
	api := r.Group("/api")
	api.GET("/users", handler)
}

func handler(c *gin.Context) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "main.go", src, 0)
	eps := scanFile(file, "main.go", fset)
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if eps[0].Path != "/api/users" {
		t.Errorf("expected path /api/users, got %q", eps[0].Path)
	}
}
