//ff:func feature=scan type=test control=sequence topic=hono
//ff:what openapiCall 테스트 헬퍼
package hono

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func openapiCall(t *testing.T, src string) (*fileInfo, *sitter.Node) {
	t.Helper()
	fi := mustParse(t, []byte(src+"\n"))

	call := findAllByType(fi.Root, "call_expression")[0]
	return fi, call
}
