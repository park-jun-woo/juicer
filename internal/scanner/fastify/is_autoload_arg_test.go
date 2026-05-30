//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what isAutoloadArg 테스트
package fastify

import "testing"

func TestIsAutoloadArg(t *testing.T) {
	names := map[string]bool{"autoload": true}

	fi := mustParse(t, []byte("register(autoload);\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	id := findChildByType(args, "identifier")
	if !isAutoloadArg(id, fi.Src, names) {
		t.Error("expected autoload identifier to match")
	}

	fi2 := mustParse(t, []byte("register(somethingElse);\n"))
	args2 := findAllByType(fi2.Root, "arguments")[0]
	other := findChildByType(args2, "identifier")
	if isAutoloadArg(other, fi2.Src, names) {
		t.Error("non-registered identifier should not match")
	}

	// non-identifier node
	str, src3 := firstNodeOfType(t, `foo("autoload");`+"\n", "string")
	if isAutoloadArg(str, src3, names) {
		t.Error("string node should not match")
	}
}
