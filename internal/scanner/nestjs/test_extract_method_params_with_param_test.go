//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractMethodParams_WithParam 테스트
package nestjs

import "testing"

func TestExtractMethodParams_WithParam(t *testing.T) {
	src := []byte(`
class C {
  findOne(@Param('id') id: string) {}
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
	result := extractMethodParams(methods[0], src)
	if len(result.pathParams) != 1 {
		t.Fatalf("expected 1 path param, got %d", len(result.pathParams))
	}
}
