//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestWalkStmts_IfElse — if-else 분기 양쪽 라우트 감지 테스트
package gogin

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestWalkStmts_IfElse(t *testing.T) {
	src := `package main

import "github.com/gin-gonic/gin"

func setup() {
	r := gin.Default()

	if true {
		r.GET("/then-branch", handler)
	} else {
		r.POST("/else-branch", handler)
	}
}

func handler(c *gin.Context) {}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	eps, _ := scanFile(file, "main.go", fset)
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}

	found := map[string]bool{}
	for _, ep := range eps {
		found[ep.Path] = true
	}
	if !found["/then-branch"] {
		t.Error("missing /then-branch")
	}
	if !found["/else-branch"] {
		t.Error("missing /else-branch")
	}
}
