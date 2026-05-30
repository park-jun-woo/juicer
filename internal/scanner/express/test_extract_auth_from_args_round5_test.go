//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what TestExtractAuthFromArgs_Round5 테스트
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestExtractAuthFromArgs_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`r.get('/x', requireAuth, handler);`))
	call := firstCallExpr(t, fi)
	args := findChildByType(call, "arguments")
	argNodes := childrenOfType(args, "")
	_ = argNodes
	// pass the argument identifiers
	var ids []*sitter.Node
	for i := 0; i < int(args.ChildCount()); i++ {
		c := args.Child(i)
		if c.Type() == "identifier" || c.Type() == "call_expression" {
			ids = append(ids, c)
		}
	}
	level, roles := extractAuthFromArgs(ids, fi.Src)
	_ = level
	_ = roles

	l2, r2 := extractAuthFromMiddlewareNodes(ids, fi.Src)
	_ = l2
	_ = r2
}
