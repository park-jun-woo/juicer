//ff:func feature=scan type=extract control=sequence
//ff:what TestScanFile_NoGin 테스트
package scanner

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestScanFile_NoGin(t *testing.T) {
	src := `package main
func main() {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "main.go", src, 0)
	eps, _ := scanFile(file, "main.go", fset)
	if len(eps) != 0 {
		t.Errorf("expected 0, got %d", len(eps))
	}
}
