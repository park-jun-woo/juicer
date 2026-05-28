//ff:func feature=scan type=parse control=sequence topic=quarkus
//ff:what Java 소스를 tree-sitter로 파싱하여 AST 루트 노드를 반환한다
package quarkus

import (
	"context"
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
	java "github.com/smacker/go-tree-sitter/java"
)

func parseJava(src []byte) (*sitter.Node, error) {
	parser := sitter.NewParser()
	parser.SetLanguage(java.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return nil, fmt.Errorf("tree-sitter parse: %w", err)
	}
	return tree.RootNode(), nil
}
