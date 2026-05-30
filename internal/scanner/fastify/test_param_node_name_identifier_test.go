//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestParamNodeName_Identifier 테스트
package fastify

import "testing"

func TestParamNodeName_Identifier(t *testing.T) {

	fi := mustParse(t, []byte("foo(myParam);\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	id := findChildByType(args, "identifier")
	if id == nil {
		t.Fatal("no identifier")
	}
	if got := paramNodeName(id, fi.Src); got != "myParam" {
		t.Fatalf("identifier: got %q", got)
	}
}
