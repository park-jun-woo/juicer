//ff:func feature=scan type=test control=sequence topic=zod
//ff:what parseTS 테스트 헬퍼
package zod

import (
	"context"
	sitter "github.com/smacker/go-tree-sitter"
	typescript "github.com/smacker/go-tree-sitter/typescript/typescript"
	"testing"
)

func parseTS(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	parser := sitter.NewParser()
	parser.SetLanguage(typescript.GetLanguage())
	tree, err := parser.ParseCtx(context.Background(), nil, b)
	if err != nil {
		t.Fatal(err)
	}
	return tree.RootNode(), b
}
