//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectPrecedingSiblingDecorators_CommentSkip 테스트 (데코레이터-메서드 사이 주석)
package nestjs

import "testing"

// A comment between the decorator and the method must not break decorator
// collection — the route should still be recognised (BUG-004).
func TestCollectPrecedingSiblingDecorators_CommentSkip(t *testing.T) {
	src := []byte(`
class LibraryController {
  @Post(':id/validate')
  // TODO: remove this once validated
  validate() {}
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	methods := findAllByType(root, "method_definition")
	if len(methods) == 0 {
		t.Fatal("no methods")
	}
	m := methods[0]
	decs := collectPrecedingSiblingDecorators(m.Parent(), m, src)
	if len(decs) != 1 {
		t.Fatalf("expected 1 decorator across comment, got %d", len(decs))
	}
	if decs[0].name != "Post" {
		t.Fatalf("expected Post decorator, got %q", decs[0].name)
	}

	// End-to-end: the route must be extracted.
	cls := findAllByType(root, "class_declaration")[0]
	eps := extractMethods(cls, src, "library.controller.ts")
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if eps[0].method != "POST" || eps[0].path != ":id/validate" {
		t.Fatalf("unexpected endpoint: %+v", eps[0])
	}
}
