//ff:func feature=scan type=test control=iteration dimension=1 topic=quarkus
//ff:what TestExtractIntLiteralFromArgList 테스트
package quarkus

import "testing"

func TestExtractIntLiteralFromArgList(t *testing.T) {
	root, _ := parseJava([]byte(`class C { void m() { status(201); } }`))
	src := []byte(`class C { void m() { status(201); } }`)
	argLists := findAllByType(root, "argument_list")
	if len(argLists) == 0 {
		t.Skip("no arg list")
	}
	for _, al := range argLists {
		if got := extractIntLiteralFromArgList(al, src); got == "201" {
			return
		}
	}
	t.Fatal("did not find 201")
}
