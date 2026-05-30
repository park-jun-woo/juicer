//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseSerializerFieldAssignment_Round5 테스트
package django

import "testing"

func TestParseSerializerFieldAssignment_Round5(t *testing.T) {
	src := []byte("name = serializers.CharField(max_length=10)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := djFirst(t, root, "expression_statement")
	f := parseSerializerFieldAssignment(stmt, src)
	if f == nil || f.Name != "name" {
		t.Fatalf("field: %+v", f)
	}
}
