//ff:func feature=scan type=parse control=sequence topic=dotnet
//ff:what C# 소스를 tree-sitter로 파싱하여 AST 루트 노드를 반환한다
package dotnet

import (
	"context"
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
	csharp "github.com/smacker/go-tree-sitter/csharp"
)

func parseCSharp(src []byte) (*sitter.Node, error) {
	parser := sitter.NewParser()
	parser.SetLanguage(csharp.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return nil, fmt.Errorf("tree-sitter parse: %w", err)
	}
	return tree.RootNode(), nil
}
