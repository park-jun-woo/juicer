//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneZodValidator_NotCall 테스트
package hono

import "testing"

func TestExtractOneZodValidator_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`const x = foo;`))
	id := findAllByType(fi.Root, "identifier")[0]
	if v := extractOneZodValidator(id, fi.Src); v != nil {
		t.Fatalf("expected nil, got %+v", v)
	}
}
