//ff:func feature=scan type=test topic=flask control=sequence
//ff:what flask 테스트 헬퍼 — Python 소스 파싱 후 첫 subscript 노드 반환
package flask

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstSubscript(t *testing.T, code string) (*sitter.Node, []byte) {
	t.Helper()
	src := []byte(code)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	subs := findAllByType(root, "subscript")
	if len(subs) == 0 {
		t.Fatalf("no subscript in %q", code)
	}
	return subs[0], src
}
