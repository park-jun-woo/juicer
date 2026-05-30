//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractRouteChain: 유효 체인 추출 / 비체인 nil 분기
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

// outermostCall returns the call_expression that is not the object of another call.
func outermostCall(fi *fileInfo) *sitter.Node {
	calls := findAllByType(fi.Root, "call_expression")
	for _, c := range calls {
		// a call whose parent chain doesn't make it an inner object
		p := c.Parent()
		isInner := false
		for p != nil {
			if p.Type() == "call_expression" {
				isInner = true
				break
			}
			if p.Type() == "expression_statement" || p.Type() == "program" {
				break
			}
			p = p.Parent()
		}
		if !isInner {
			return c
		}
	}
	if len(calls) > 0 {
		return calls[0]
	}
	return nil
}

func TestExtractRouteChain_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/:id').get(getH).put(putH);`))
	call := outermostCall(fi)
	routes := extractRouteChain(call, fi.Src, map[string]bool{"router": true})
	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %+v", routes)
	}
	if routes[0].Path != "/:id" || routes[1].Path != "/:id" {
		t.Fatalf("unexpected paths %+v", routes)
	}
}

func TestExtractRouteChain_NotChain(t *testing.T) {
	fi := mustParse(t, []byte(`foo();`))
	var first *sitter.Node = firstCallExpr(t, fi)
	if routes := extractRouteChain(first, fi.Src, map[string]bool{"router": true}); routes != nil {
		t.Fatalf("expected nil, got %+v", routes)
	}
}
