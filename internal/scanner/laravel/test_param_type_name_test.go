//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestParamTypeName 테스트
package laravel

import "testing"

func TestParamTypeName(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m(int $id, FooRequest $r, $untyped) {} }`)
	sps := findAllByType(fi.root, "simple_parameter")
	if len(sps) < 3 {
		t.Fatalf("expected 3 params, got %d", len(sps))
	}
	if got := paramTypeName(sps[0], fi.src); got != "int" {
		t.Fatalf("primitive: %q", got)
	}
	if got := paramTypeName(sps[1], fi.src); got != "FooRequest" {
		t.Fatalf("named: %q", got)
	}
	if got := paramTypeName(sps[2], fi.src); got != "" {
		t.Fatalf("untyped: %q", got)
	}
}
