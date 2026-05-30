//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractPluginRef_Identifier 테스트
package fastify

import "testing"

func TestExtractPluginRef_Identifier(t *testing.T) {
	fi := mustParse(t, []byte("foo(myPlugin);\n"))

	args := findAllByType(fi.Root, "arguments")[0]
	id := findChildByType(args, "identifier")
	if id == nil {
		t.Fatal("no identifier arg")
	}
	if got := extractPluginRef(id, fi.Src); got != "myPlugin" {
		t.Fatalf("identifier: got %q", got)
	}
}
