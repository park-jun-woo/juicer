//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractAssignedValue_Identifier 테스트
package django

import "testing"

func TestExtractAssignedValue_Identifier(t *testing.T) {
	src := []byte("serializer_class = UserSerializer\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := firstExprStatement(root)
	if stmt == nil {
		t.Fatal("no expression_statement")
	}
	got := extractAssignedValue(stmt, "serializer_class", src)
	if got != "UserSerializer" {
		t.Fatalf("got %q, want UserSerializer", got)
	}
}
