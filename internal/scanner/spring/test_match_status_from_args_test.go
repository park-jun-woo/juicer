//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestMatchStatusFromArgs 테스트
package spring

import "testing"

func TestMatchStatusFromArgs(t *testing.T) {
	root, src := parseS(t, `class C { void m() { ResponseEntity.status(HttpStatus.CREATED).build(); } }`)
	argLists := findAllByType(root, "argument_list")
	for _, al := range argLists {
		if code := extractStatusFromArgList(al, src); code == "201" {
			return
		}
	}
	t.Fatal("did not find 201")
}
