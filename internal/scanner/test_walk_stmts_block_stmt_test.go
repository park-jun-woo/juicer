//ff:func feature=scan type=extract control=sequence
//ff:what TestWalkStmts_BlockStmt 테스트
package scanner

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestWalkStmts_BlockStmt(t *testing.T) {
	// Block statement containing route calls
	src := `package main

import "github.com/gin-gonic/gin"

func setup() {
	r := gin.Default()
	{
		r.GET("/nested", handler)
	}
}

func handler(c *gin.Context) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "main.go", src, 0)
	eps := scanFile(file, "main.go", fset)
	if len(eps) != 1 {
		t.Errorf("expected 1 endpoint from nested block, got %d", len(eps))
	}
}
