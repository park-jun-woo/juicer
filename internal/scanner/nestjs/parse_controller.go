//ff:func feature=scan type=parse control=sequence topic=nestjs
//ff:what TypeScript 소스를 tree-sitter로 파싱하여 AST 루트 노드를 반환한다
package nestjs

import (
	"context"
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
	typescript "github.com/smacker/go-tree-sitter/typescript/typescript"
)

// parseTypeScript parses TypeScript source bytes and returns the root AST node.
func parseTypeScript(src []byte) (*sitter.Node, error) {
	parser := sitter.NewParser()
	parser.SetLanguage(typescript.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return nil, fmt.Errorf("tree-sitter parse: %w", err)
	}
	return tree.RootNode(), nil
}
