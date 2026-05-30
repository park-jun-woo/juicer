//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractAssignedValue_NoAssignment 테스트
package django

import "testing"

func TestExtractAssignedValue_NoAssignment(t *testing.T) {
	src := []byte("foo()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := firstExprStatement(root)
	if got := extractAssignedValue(stmt, "x", src); got != "" {
		t.Fatalf("expected empty for non-assignment, got %q", got)
	}
}
