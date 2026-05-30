//ff:func feature=scan type=test control=sequence topic=express
//ff:what firstObject 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstObject(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	objs := findAllByType(fi.Root, "object")
	if len(objs) == 0 {
		t.Fatal("no object")
	}
	return objs[0]
}
