//ff:func feature=scan type=test topic=fastify control=sequence
//ff:what applyTypeBoxOptions format/numeric/enum 옵션 반영 및 비-object 무시 테스트
package fastify

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyTypeBoxOptions(t *testing.T) {
	// format option
	fi := mustParse(t, []byte(`Type.String({ format: 'email' })`))
	call := findAllByType(fi.Root, "call_expression")[0]
	f := scanner.Field{Type: "string"}
	applyTypeBoxOptions(&f, call, fi.Src)
	if f.Type != "email" {
		t.Errorf("format should set type to email: %+v", f)
	}

	// no options object -> unchanged
	fi2 := mustParse(t, []byte(`Type.String()`))
	c2 := findAllByType(fi2.Root, "call_expression")[0]
	g := scanner.Field{Type: "string"}
	applyTypeBoxOptions(&g, c2, fi2.Src)
	if g.Type != "string" {
		t.Errorf("no opts should be unchanged: %+v", g)
	}
}
