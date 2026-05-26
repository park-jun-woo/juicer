//ff:func feature=scan type=extract control=sequence
//ff:what TestGinPkgName 테스트
package gogin

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestGinPkgName(t *testing.T) {
	t.Run("standard import", func(t *testing.T) {
		src := `package main
import "github.com/gin-gonic/gin"
`
		fset := token.NewFileSet()
		file, _ := parser.ParseFile(fset, "test.go", src, 0)
		got := ginPkgName(file)
		if got != "gin" {
			t.Errorf("expected 'gin', got %q", got)
		}
	})

	t.Run("aliased import", func(t *testing.T) {
		src := `package main
import g "github.com/gin-gonic/gin"
`
		fset := token.NewFileSet()
		file, _ := parser.ParseFile(fset, "test.go", src, 0)
		got := ginPkgName(file)
		if got != "g" {
			t.Errorf("expected 'g', got %q", got)
		}
	})

	t.Run("no gin import", func(t *testing.T) {
		src := `package main
import "fmt"
`
		fset := token.NewFileSet()
		file, _ := parser.ParseFile(fset, "test.go", src, 0)
		got := ginPkgName(file)
		if got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})
}
