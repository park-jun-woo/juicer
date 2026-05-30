//ff:func feature=scan type=test control=sequence topic=express
//ff:what findInitValue: call_expression / new_expression / none 분기
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstDeclarator(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	ds := findAllByType(fi.Root, "variable_declarator")
	if len(ds) == 0 {
		t.Fatal("no variable_declarator")
	}
	return ds[0]
}

func TestFindInitValue_Call(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express.Router();`))
	if v := findInitValue(firstDeclarator(t, fi)); v == nil || v.Type() != "call_expression" {
		t.Fatalf("got %v", v)
	}
}

func TestFindInitValue_New(t *testing.T) {
	fi := mustParse(t, []byte(`const r = new Router();`))
	if v := findInitValue(firstDeclarator(t, fi)); v == nil || v.Type() != "new_expression" {
		t.Fatalf("got %v", v)
	}
}

func TestFindInitValue_None(t *testing.T) {
	fi := mustParse(t, []byte(`const r = 42;`))
	if v := findInitValue(firstDeclarator(t, fi)); v != nil {
		t.Fatalf("expected nil, got %v", v.Type())
	}
}
