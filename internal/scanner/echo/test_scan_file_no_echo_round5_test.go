//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestScanFile_NoEcho_Round5 테스트
package echo

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestScanFile_NoEcho_Round5(t *testing.T) {
	src := `package m
func F() {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	eps, hmap := scanFile(nil, file, "m.go", fset)
	if eps != nil || hmap != nil {
		t.Fatalf("no echo import -> nil, got %v %v", eps, hmap)
	}
}
