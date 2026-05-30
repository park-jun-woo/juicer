//ff:func feature=scan type=test control=sequence
//ff:what resolveEmbedded — 임베딩 타입 필드 추출 테스트
package fiber

import (
	"go/types"
	"testing"
)

func TestResolveEmbedded_Struct(t *testing.T) {
	src := `package m
type Base struct {
	ID int ` + "`json:\"id\"`" + `
}
type User struct {
	Base
}
`
	st, _ := structFields(t, src, "User")
	// the embedded field's type is *types.Named "Base"
	embType := st.Field(0).Type()
	fields := resolveEmbedded(embType, map[string]bool{})
	if len(fields) != 1 || fields[0].Name != "ID" {
		t.Fatalf("expected ID field, got %v", fields)
	}
}

func TestResolveEmbedded_Cycle(t *testing.T) {
	src := `package m
type T struct { X int }
`
	st, _ := structFields(t, src, "T")
	typ := st.Field(0).Type() // int
	// mark its key as visited -> returns nil
	visited := map[string]bool{typ.String(): true}
	if got := resolveEmbedded(typ, visited); got != nil {
		t.Fatalf("expected nil for visited, got %v", got)
	}
}

func TestResolveEmbedded_NotStruct(t *testing.T) {
	// a basic int type -> not a struct -> nil
	if got := resolveEmbedded(types.Typ[types.Int], map[string]bool{}); got != nil {
		t.Fatalf("expected nil for non-struct, got %v", got)
	}
}
