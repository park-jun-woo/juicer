//ff:func feature=scan type=test control=sequence topic=actix
//ff:what extractResponses — 함수 본문에서 응답 추출, block 없는 경우 nil
package actix

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"

	sitter "github.com/smacker/go-tree-sitter"
)

// findFuncByName returns the first function_item node whose name matches.
func findFuncByName(root *sitter.Node, src []byte, name string) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "function_item" || n.Type() == "function_signature_item" {
			id := findChildByType(n, "identifier")
			if id != nil && nodeText(id, src) == name {
				found = n
				return
			}
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return found
}

func TestExtractResponses_WithBlock(t *testing.T) {
	src := []byte(`
struct CreatedURL { a: String }
async fn create_url() -> HttpResponse {
    HttpResponse::Created().json(CreatedURL{a: x})
}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fn := findFuncByName(root, src, "create_url")
	if fn == nil {
		t.Fatal("function not found")
	}
	idx := buildStructIndex([]*fileInfo{{src: src, root: root}})
	cache := map[string][]scanner.Field{}
	resps := extractResponses(fn, src, idx, cache)
	if len(resps) != 1 {
		t.Fatalf("expected 1 response, got %d: %+v", len(resps), resps)
	}
	if resps[0].Status != "201" {
		t.Errorf("status = %q, want 201", resps[0].Status)
	}
}

func TestExtractResponses_NoBlock(t *testing.T) {
	// A function signature without a body block (trait method declaration)
	// yields no responses.
	src := []byte(`
trait T {
    async fn handler() -> HttpResponse;
}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fn := findFuncByName(root, src, "handler")
	if fn == nil {
		t.Fatal("function signature not found")
	}
	idx := buildStructIndex([]*fileInfo{{src: src, root: root}})
	if resps := extractResponses(fn, src, idx, map[string][]scanner.Field{}); resps != nil {
		t.Fatalf("expected nil responses for block-less function, got %+v", resps)
	}
}
