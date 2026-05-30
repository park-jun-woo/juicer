//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestResolveNestedFields_And_Embedded_Round5 테스트
package echo

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_And_Embedded_Round5(t *testing.T) {
	_, info := checkSrc(t, dtoSrc)
	typ := namedType(t, info, "U")
	nested := resolveNestedFields(typ, map[string]bool{})
	if len(nested) == 0 {
		t.Fatalf("expected nested fields, got %+v", nested)
	}

	named := typ.(*types.Named)
	emb := resolveEmbedded(named, map[string]bool{})
	if len(emb) == 0 {
		t.Fatalf("expected embedded resolution, got %+v", emb)
	}
}
