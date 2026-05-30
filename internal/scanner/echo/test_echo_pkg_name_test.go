//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestEchoPkgName 테스트
package echo

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestEchoPkgName(t *testing.T) {
	src := `package m
import "github.com/labstack/echo/v4"
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	if got := echoPkgName(file); got != "echo" {
		t.Fatalf("got %q", got)
	}

	src2 := `package m
import e "github.com/labstack/echo/v4"
`
	file2, _ := parser.ParseFile(fset, "m2.go", src2, 0)
	if got := echoPkgName(file2); got != "e" {
		t.Fatalf("aliased: %q", got)
	}

	src3 := `package m
import "fmt"
`
	file3, _ := parser.ParseFile(fset, "m3.go", src3, 0)
	if got := echoPkgName(file3); got != "" {
		t.Fatalf("no echo import: %q", got)
	}
}
