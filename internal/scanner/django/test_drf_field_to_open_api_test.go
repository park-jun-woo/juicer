//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what DRF 필드 타입을 OpenAPI 타입으로 변환한다
package django

import "testing"

func TestDRFFieldToOpenAPI(t *testing.T) {
	tests := []struct {
		field  string
		wantT  string
		wantF  string
	}{
		{"CharField", "string", ""},
		{"EmailField", "string", "email"},
		{"IntegerField", "integer", ""},
		{"FloatField", "number", "float"},
		{"BooleanField", "boolean", ""},
		{"DateTimeField", "string", "date-time"},
		{"DateField", "string", "date"},
		{"UUIDField", "string", "uuid"},
		{"URLField", "string", "uri"},
		{"FileField", "string", "binary"},
		{"ListField", "array", ""},
		{"JSONField", "object", ""},
		{"PrimaryKeyRelatedField", "integer", ""},
	}

	for _, tt := range tests {
		got := drfFieldToOpenAPI(tt.field)
		if got.Type != tt.wantT || got.Format != tt.wantF {
			t.Errorf("drfFieldToOpenAPI(%q) = {%q, %q}, want {%q, %q}",
				tt.field, got.Type, got.Format, tt.wantT, tt.wantF)
		}
	}
}
