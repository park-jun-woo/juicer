//ff:func feature=scan type=test control=sequence topic=express
//ff:what firstArrow 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstArrow(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	a := findAllByType(fi.Root, "arrow_function")
	if len(a) == 0 {
		t.Fatal("no arrow_function")
	}
	return a[0]
}
