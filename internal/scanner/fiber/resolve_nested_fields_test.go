//ff:func feature=scan type=test control=sequence
//ff:what resolveNestedFields — 중첩 필드 추출 테스트
package fiber

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_NamedStruct(t *testing.T) {
	src := `package m
type Inner struct { Z int ` + "`json:\"z\"`" + ` }
type Outer struct {
	In    Inner
	List  []Inner
	When  timeStub
}
type timeStub struct{ T int }
`
	st, _ := structFields(t, src, "Outer")
	// In Inner -> nested fields
	in := resolveNestedFields(st.Field(0).Type(), map[string]bool{})
	if len(in) != 1 || in[0].Name != "Z" {
		t.Fatalf("named struct nested = %v", in)
	}
	// []Inner -> element struct
	list := resolveNestedFields(st.Field(1).Type(), map[string]bool{})
	if len(list) != 1 || list[0].Name != "Z" {
		t.Fatalf("slice of struct nested = %v", list)
	}
}

func TestResolveNestedFields_ArrayOfStruct(t *testing.T) {
	src := `package m
type E struct { N int ` + "`json:\"n\"`" + ` }
type A struct { Items [3]E }
`
	st, _ := structFields(t, src, "A")
	got := resolveNestedFields(st.Field(0).Type(), map[string]bool{})
	if len(got) != 1 || got[0].Name != "N" {
		t.Fatalf("array of struct nested = %v", got)
	}
}

func TestResolveNestedFields_WellKnown(t *testing.T) {
	src := `package m
import "time"
type T struct { When time.Time }
`
	st, _ := structFields(t, src, "T")
	if got := resolveNestedFields(st.Field(0).Type(), map[string]bool{}); got != nil {
		t.Fatalf("time.Time should not expand, got %v", got)
	}
}

func TestResolveNestedFields_NonStruct(t *testing.T) {
	if got := resolveNestedFields(types.Typ[types.Int], map[string]bool{}); got != nil {
		t.Fatalf("int should yield nil, got %v", got)
	}
}

func TestResolveNestedFields_Cycle(t *testing.T) {
	src := `package m
type S struct { X int }
`
	st, _ := structFields(t, src, "S")
	typ := st.Field(0).Type()
	visited := map[string]bool{typ.String(): true}
	if got := resolveNestedFields(typ, visited); got != nil {
		t.Fatalf("visited should yield nil, got %v", got)
	}
}
