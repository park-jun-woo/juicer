//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestCollectChainMethodsRecursive_NilAndDefault_Round5 테스트
package zod

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestCollectChainMethodsRecursive_NilAndDefault_Round5(t *testing.T) {
	var methods []ChainMethod

	collectChainMethodsRecursive(nil, nil, &methods)
	if len(methods) != 0 {
		t.Fatal("nil node should not append")
	}

	root, src := parseTS(t, "const x = 1;")
	var idNode *sitter.Node
	walkNodes(root, func(n *sitter.Node) {
		if idNode == nil && n.Type() == "identifier" {
			idNode = n
		}
	})
	if idNode == nil {
		t.Fatal("no identifier")
	}
	collectChainMethodsRecursive(idNode, src, &methods)
	if len(methods) != 0 {
		t.Fatal("identifier should be no-op")
	}
}
