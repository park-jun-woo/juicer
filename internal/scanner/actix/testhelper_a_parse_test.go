//ff:func feature=scan type=test control=sequence topic=actix
//ff:what aParse 테스트 헬퍼
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func aParse(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parseRust(b)
	if err != nil {
		t.Fatal(err)
	}
	return root, b
}
