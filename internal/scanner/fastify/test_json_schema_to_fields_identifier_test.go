//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestJSONSchemaToFields_Identifier 테스트
package fastify

import "testing"

func TestJSONSchemaToFields_Identifier(t *testing.T) {

	fi := mustParse(t, []byte("const x = SomeSchema;\n"))
	ids := findAllByType(fi.Root, "identifier")
	var ref = ids[len(ids)-1] // the RHS reference
	if got := jsonSchemaToFields(ref, fi.Src); got != nil {
		t.Fatalf("expected nil for identifier, got %v", got)
	}
}
