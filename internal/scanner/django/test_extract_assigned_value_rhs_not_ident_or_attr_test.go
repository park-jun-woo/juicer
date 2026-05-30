//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractAssignedValue_RHSNotIdentOrAttr 테스트
package django

import "testing"

func TestExtractAssignedValue_RHSNotIdentOrAttr(t *testing.T) {

	src := []byte("serializer_class = 42\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := firstExprStatement(root)
	if got := extractAssignedValue(stmt, "serializer_class", src); got != "" {
		t.Fatalf("expected empty for literal RHS, got %q", got)
	}
}
