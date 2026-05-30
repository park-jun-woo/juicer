//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what TestExtractHandlerAndMiddleware_Round5 테스트
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestExtractHandlerAndMiddleware_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`r.get('/x', mw, handler);`))
	call := firstCallExpr(t, fi)
	args := findChildByType(call, "arguments")
	var argNodes []*sitter.Node
	for i := 0; i < int(args.ChildCount()); i++ {
		c := args.Child(i)
		switch c.Type() {
		case "(", ")", ",":
		default:
			argNodes = append(argNodes, c)
		}
	}

	handler, mw := extractHandlerAndMiddleware(argNodes, fi.Src)
	if handler != "handler" {
		t.Fatalf("handler: %q", handler)
	}
	if len(mw) != 1 || mw[0] != "mw" {
		t.Fatalf("middleware: %v", mw)
	}
}
