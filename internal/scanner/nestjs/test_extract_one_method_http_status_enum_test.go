//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractOneMethod_HttpStatusEnum 테스트
package nestjs

import "testing"

func TestExtractOneMethod_HttpStatusEnum(t *testing.T) {
	src := []byte(`
class C {
  @HttpCode(HttpStatus.OK)
  @Post('login')
  login() {}
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
	if ep.statusCode != 200 {
		t.Fatalf("expected statusCode=200, got %d", ep.statusCode)
	}
}
