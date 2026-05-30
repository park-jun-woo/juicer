//ff:func feature=scan type=test control=sequence topic=express
//ff:what firstArray 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstArray(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	arrs := findAllByType(fi.Root, "array")
	if len(arrs) == 0 {
		t.Fatal("no array")
	}
	return arrs[0]
}
