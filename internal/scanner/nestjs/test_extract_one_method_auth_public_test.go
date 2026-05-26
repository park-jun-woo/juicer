//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractOneMethod_AuthLevel_Public 테스트
package nestjs

import "testing"

func TestExtractOneMethod_AuthLevel_Public(t *testing.T) {
	src := []byte(`
class C {
  @Public()
  @Get('/')
  home() {}
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	methods := findAllByType(root, "method_definition")
	if len(methods) == 0 {
		t.Fatal("no methods found")
	}
	ep, ok := extractOneMethod(methods[0], src, "test.ts")
	if !ok {
		t.Fatal("expected ok")
	}
	if ep.authLevel != "public" {
		t.Fatalf("expected public, got %q", ep.authLevel)
	}
}
