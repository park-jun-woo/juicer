//ff:func feature=scan type=test control=sequence topic=express
//ff:what matchAnyDecl: 함수선언 / 변수선언 / export선언 / 미매칭 분기
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func topChild(t *testing.T, fi *fileInfo, typ string) *sitter.Node {
	t.Helper()
	root := fi.Root
	for i := 0; i < int(root.ChildCount()); i++ {
		if root.Child(i).Type() == typ {
			return root.Child(i)
		}
	}
	t.Fatalf("no top-level %s", typ)
	return nil
}

func TestMatchAnyDecl_Function(t *testing.T) {
	fi := mustParse(t, []byte(`function handler(req, res) { res.json({}); }`))
	if body := matchAnyDecl(topChild(t, fi, "function_declaration"), fi.Src, "handler"); body == nil {
		t.Fatal("expected function body")
	}
}

func TestMatchAnyDecl_Variable(t *testing.T) {
	fi := mustParse(t, []byte(`const handler = (req, res) => { res.json({}); };`))
	if body := matchAnyDecl(topChild(t, fi, "lexical_declaration"), fi.Src, "handler"); body == nil {
		t.Fatal("expected variable body")
	}
}

func TestMatchAnyDecl_Exported(t *testing.T) {
	fi := mustParse(t, []byte(`export function handler(req, res) { res.json({}); }`))
	if body := matchAnyDecl(topChild(t, fi, "export_statement"), fi.Src, "handler"); body == nil {
		t.Fatal("expected exported body")
	}
}

func TestMatchAnyDecl_NoMatch(t *testing.T) {
	fi := mustParse(t, []byte(`function handler() {}`))
	if body := matchAnyDecl(topChild(t, fi, "function_declaration"), fi.Src, "other"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
