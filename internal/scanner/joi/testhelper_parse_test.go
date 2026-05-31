//ff:func feature=scan type=test topic=joi control=sequence
//ff:what joi 테스트용 TypeScript 파서 헬퍼 — 소스를 파싱해 루트 노드/소스 반환
package joi

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	typescript "github.com/smacker/go-tree-sitter/typescript/typescript"
)

func parseJoiTS(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	parser := sitter.NewParser()
	parser.SetLanguage(typescript.GetLanguage())
	b := []byte(src)
	tree, err := parser.ParseCtx(context.Background(), nil, b)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	return tree.RootNode(), b
}
