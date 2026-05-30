//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestMatchBodyConstructors 테스트
package spring

import "testing"

func TestMatchBodyConstructors(t *testing.T) {
	root, src := parseS(t, `class C { void m() { return new ResponseEntity<>(HttpStatus.OK); } }`)
	body := findAllByType(root, "block")[0]
	_ = matchBodyConstructors(body, src)
}
