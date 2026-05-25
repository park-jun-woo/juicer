//ff:func feature=scan type=extract control=sequence
//ff:what TestScanFile_WithGinRoutes 테스트
package scanner

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestScanFile_WithGinRoutes(t *testing.T) {
	src := `package main

import "github.com/gin-gonic/gin"

func setup() {
	r := gin.Default()
	r.GET("/users", handler)
	r.POST("/users", createHandler)
}

func handler(c *gin.Context) {}
func createHandler(c *gin.Context) {}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	eps := scanFile(file, "main.go", fset)
	if len(eps) != 2 {
		t.Errorf("expected 2 endpoints, got %d", len(eps))
	}
}
