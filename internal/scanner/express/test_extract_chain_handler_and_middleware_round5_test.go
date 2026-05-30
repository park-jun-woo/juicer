//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractChainHandlerAndMiddleware_Round5 테스트
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestExtractChainHandlerAndMiddleware_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`r.route('/x').get(mw, handler);`))
	// find the .get(...) call
	var getCall *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if getCall != nil {
			return
		}
		if n.Type() == "call_expression" {
			fn := n.ChildByFieldName("function")
			if fn != nil && fn.Type() == "member_expression" {
				prop := fn.ChildByFieldName("property")
				if prop != nil && nodeText(prop, fi.Src) == "get" {
					getCall = n
				}
			}
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(fi.Root)
	if getCall == nil {
		t.Fatal("no get call")
	}
	handler, mw, _, _, _ := extractChainHandlerAndMiddleware(getCall, fi.Src)
	if handler != "handler" {
		t.Fatalf("handler: %q", handler)
	}
	if len(mw) != 1 || mw[0] != "mw" {
		t.Fatalf("middleware: %v", mw)
	}
}
