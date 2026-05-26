//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractOneMethod_Get 테스트
package nestjs

import "testing"

func TestExtractOneMethod_Get(t *testing.T) {
	src := []byte(`
class UsersController {
  @Get('all')
  findAll() {}
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
	ep, ok := extractOneMethod(methods[0], src, "test.ts")
	if !ok {
		t.Fatal("expected ok")
	}
	if ep.method != "GET" {
		t.Fatalf("expected GET, got %s", ep.method)
	}
	if ep.path != "all" {
		t.Fatalf("expected all, got %q", ep.path)
	}
}
