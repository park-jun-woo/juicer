//ff:func feature=scan type=parse control=sequence topic=fastify
//ff:what TypeScript 소스를 tree-sitter로 파싱하여 AST 루트 노드를 반환한다
package fastify

import (
	"context"
	"fmt"
	"os"

	sitter "github.com/smacker/go-tree-sitter"
	typescript "github.com/smacker/go-tree-sitter/typescript/typescript"
)

func parseFile(path string) (*fileInfo, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	parser := sitter.NewParser()
	parser.SetLanguage(typescript.GetLanguage())
	tree, err := parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return nil, fmt.Errorf("tree-sitter parse %s: %w", path, err)
	}
	return &fileInfo{Path: path, Root: tree.RootNode(), Src: src}, nil
}
