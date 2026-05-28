//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what Django URL 컨버터 타입을 OpenAPI 타입으로 변환한다
package django

import "testing"

func TestDjangoConverterToOpenAPI(t *testing.T) {
	tests := []struct {
		converter string
		wantT     string
		wantF     string
	}{
		{"int", "integer", ""},
		{"uuid", "string", "uuid"},
		{"slug", "string", ""},
		{"str", "string", ""},
		{"", "string", ""},
		{"path", "string", ""},
	}

	for _, tt := range tests {
		got := djangoConverterToOpenAPI(tt.converter)
		if got.Type != tt.wantT || got.Format != tt.wantF {
			t.Errorf("djangoConverterToOpenAPI(%q) = {%q, %q}, want {%q, %q}",
				tt.converter, got.Type, got.Format, tt.wantT, tt.wantF)
		}
	}
}
