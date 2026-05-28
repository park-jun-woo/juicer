//ff:func feature=scan type=test control=sequence topic=hono
//ff:what Zod z.number().int().min(0) 체인 파싱 테스트
package hono

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner/zod"
)

func TestParseZodChain_NumberIntMin(t *testing.T) {
	src := []byte(`const x = z.number().int().min(0)`)
	fi := mustParse(t, src)
	decls := findAllByType(fi.Root, "lexical_declaration")
	declarators := childrenOfType(decls[0], "variable_declarator")
	value := declarators[0].ChildByFieldName("value")
	f := zod.ParseChain(value, fi.Src)
	if f.Type != "integer" {
		t.Errorf("expected type integer, got %s", f.Type)
	}
	if f.Minimum == nil || *f.Minimum != 0 {
		t.Errorf("expected minimum 0")
	}
}
