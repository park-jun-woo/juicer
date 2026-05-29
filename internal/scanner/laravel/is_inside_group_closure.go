//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 노드가 root까지의 ->group(closure) 안에 있는지 보고한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// isInsideGroupClosure reports whether node sits inside a ->group(closure)
// located between node and root (exclusive). root marks the current scope:
// at the file level it is the program node; inside a group body it is that
// body, so routes directly in the body are kept while deeper-nested group
// routes are deferred to the recursive group walk.
func isInsideGroupClosure(node, root *sitter.Node, fi fileInfo) bool {
	for n := node.Parent(); n != nil && !sameNode(n, root); n = n.Parent() {
		if n.Type() != "anonymous_function_creation_expression" && n.Type() != "arrow_function" {
			continue
		}
		if isGroupCallArgument(n, fi) {
			return true
		}
	}
	return false
}
