//ff:func feature=scan type=test control=sequence
//ff:what TestScanFile_NoFiber 테스트
package fiber

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestScanFile_NoFiber(t *testing.T) {
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", "package m\nfunc F() {}\n", 0)
	eps, hmap := scanFile(file, "m.go", fset)
	if eps != nil || hmap != nil {
		t.Fatalf("non-fiber file should yield nil, got %v %v", eps, hmap)
	}
}
