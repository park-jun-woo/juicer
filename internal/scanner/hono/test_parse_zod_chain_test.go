//ff:func feature=scan type=test control=sequence topic=hono
//ff:what Zod z.string().email() 체인 파싱 테스트
package hono

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner/zod"
)

func TestParseZodChain_StringEmail(t *testing.T) {
	src := []byte(`const x = z.string().email()`)
	fi := mustParse(t, src)
	decls := findAllByType(fi.Root, "lexical_declaration")
	if len(decls) == 0 {
		t.Fatal("no declarations found")
	}
	declarators := childrenOfType(decls[0], "variable_declarator")
	if len(declarators) == 0 {
		t.Fatal("no declarators found")
	}
	value := declarators[0].ChildByFieldName("value")
	if value == nil {
		t.Fatal("no value")
	}
	f := zod.ParseChain(value, fi.Src)
	if f.Type != "string" {
		t.Errorf("expected type string, got %s", f.Type)
	}
	if f.Validate != "email" {
		t.Errorf("expected validate email, got %s", f.Validate)
	}
}
