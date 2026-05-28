//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what 테스트 헬퍼: 소스 바이트를 tree-sitter로 파싱
package supafunc

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	typescript "github.com/smacker/go-tree-sitter/typescript/typescript"
)

func mustParse(t *testing.T, src []byte) *fileInfo {
	t.Helper()
	parser := sitter.NewParser()
	parser.SetLanguage(typescript.GetLanguage())
	tree, err := parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	return &fileInfo{Path: "test.ts", Root: tree.RootNode(), Src: src}
}
