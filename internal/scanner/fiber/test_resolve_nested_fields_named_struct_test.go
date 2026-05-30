//ff:func feature=scan type=test control=sequence
//ff:what TestResolveNestedFields_NamedStruct 테스트
package fiber

import "testing"

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

	in := resolveNestedFields(st.Field(0).Type(), map[string]bool{})
	if len(in) != 1 || in[0].Name != "Z" {
		t.Fatalf("named struct nested = %v", in)
	}

	list := resolveNestedFields(st.Field(1).Type(), map[string]bool{})
	if len(list) != 1 || list[0].Name != "Z" {
		t.Fatalf("slice of struct nested = %v", list)
	}
}
