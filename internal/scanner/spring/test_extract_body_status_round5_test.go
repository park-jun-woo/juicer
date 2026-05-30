//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractBodyStatus_Round5 테스트
package spring

import "testing"

func TestExtractBodyStatus_Round5(t *testing.T) {

	root, src := sParse(t, `class C {
		public Object create() {
			return ResponseEntity.status(HttpStatus.CREATED).build();
		}
	}`)
	m := sFirst(t, root, "method_declaration")
	if got := extractBodyStatus(m, src); got == "" {
		t.Fatalf("expected a status code, got empty")
	}

	root2, src2 := sParse(t, `class C { public Object g() { return null; } }`)
	m2 := sFirst(t, root2, "method_declaration")
	if got := extractBodyStatus(m2, src2); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
