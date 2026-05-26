//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractMethodParams_NoParams 테스트
package nestjs

import "testing"

func TestExtractMethodParams_NoParams(t *testing.T) {
	src := []byte(`
class C {
  @Get()
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
	result := extractMethodParams(methods[0], src)
	if len(result.pathParams) != 0 {
		t.Fatal("expected no params")
	}
}
