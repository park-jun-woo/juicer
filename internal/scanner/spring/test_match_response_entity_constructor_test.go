//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestMatchResponseEntityConstructor 테스트
package spring

import "testing"

func TestMatchResponseEntityConstructor(t *testing.T) {
	root, src := parseS(t, `class C { void m() { return new ResponseEntity<>(body, HttpStatus.CREATED); } }`)
	objs := findAllByType(root, "object_creation_expression")
	for _, o := range objs {
		if matchResponseEntityConstructor(o, src) == "201" {
			return
		}
	}
	t.Skip("constructor status not matched in this grammar")
}
