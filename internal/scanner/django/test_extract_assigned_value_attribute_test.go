//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractAssignedValue_Attribute 테스트
package django

import "testing"

func TestExtractAssignedValue_Attribute(t *testing.T) {
	src := []byte("serializer_class = mod.UserSerializer\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := firstExprStatement(root)
	got := extractAssignedValue(stmt, "serializer_class", src)
	if got != "mod.UserSerializer" {
		t.Fatalf("got %q, want mod.UserSerializer", got)
	}
}
