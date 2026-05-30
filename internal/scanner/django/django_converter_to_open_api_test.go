//ff:func feature=scan type=test control=sequence topic=django
//ff:what djangoConverterToOpenAPI — default(미지정) 컨버터 분기를 검증
package django

import "testing"

func TestDjangoConverterToOpenAPI_Default(t *testing.T) {
	// An unknown converter falls through to the default string mapping.
	got := djangoConverterToOpenAPI("customthing")
	if got.Type != "string" || got.Format != "" {
		t.Fatalf("default = {%q, %q}, want {string, }", got.Type, got.Format)
	}
}
