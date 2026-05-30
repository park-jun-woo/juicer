//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractAssignedValue_WrongName 테스트
package django

import "testing"

func TestExtractAssignedValue_WrongName(t *testing.T) {
	src := []byte("other = Value\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := firstExprStatement(root)
	if got := extractAssignedValue(stmt, "serializer_class", src); got != "" {
		t.Fatalf("expected empty for wrong LHS name, got %q", got)
	}
}
