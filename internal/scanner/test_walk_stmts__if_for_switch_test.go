//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestWalkStmts_IfForSwitch — if/for/switch 내부 라우트 감지 테스트
package scanner

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestWalkStmts_IfForSwitch(t *testing.T) {
	src := `package main

import "github.com/gin-gonic/gin"

func setup() {
	r := gin.Default()

	if true {
		r.GET("/in-if", handler)
	}

	for i := 0; i < 1; i++ {
		r.POST("/in-for", handler)
	}

	switch "a" {
	case "a":
		r.PUT("/in-switch", handler)
	}
}

func handler(c *gin.Context) {}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	eps := scanFile(file, "main.go", fset)
	if len(eps) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(eps))
	}

	expected := map[string]string{
		"/in-if":     "GET",
		"/in-for":    "POST",
		"/in-switch": "PUT",
	}
	for _, ep := range eps {
		wantMethod, ok := expected[ep.Path]
		if !ok {
			t.Errorf("unexpected endpoint path: %s", ep.Path)
			continue
		}
		if ep.Method != wantMethod {
			t.Errorf("path %s: expected method %s, got %s", ep.Path, wantMethod, ep.Method)
		}
	}
}
