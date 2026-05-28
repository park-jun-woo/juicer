//ff:func feature=scan type=parse control=sequence topic=django
//ff:what Python 소스를 tree-sitter로 파싱하여 AST 루트 노드를 반환한다
package django

import (
	"context"
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
	python "github.com/smacker/go-tree-sitter/python"
)

// parsePython parses Python source bytes and returns the root AST node.
func parsePython(src []byte) (*sitter.Node, error) {
	parser := sitter.NewParser()
	parser.SetLanguage(python.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return nil, fmt.Errorf("tree-sitter parse: %w", err)
	}
	return tree.RootNode(), nil
}
