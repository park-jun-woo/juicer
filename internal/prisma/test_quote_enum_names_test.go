//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what quoteEnumNames enum 타입명 인용 복사 테스트
package prisma

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/ddl"
)

func TestQuoteEnumNames(t *testing.T) {
	in := []ddl.EnumType{{Name: "Role", Values: []string{"ADMIN", "USER"}}}
	out := quoteEnumNames(in)
	if len(out) != 1 || out[0].Name != `"Role"` {
		t.Fatalf("got %+v", out)
	}
	if len(out[0].Values) != 2 || out[0].Values[0] != "ADMIN" {
		t.Errorf("values changed: %+v", out[0].Values)
	}
	// original untouched
	if in[0].Name != "Role" {
		t.Errorf("original mutated: %q", in[0].Name)
	}
}
