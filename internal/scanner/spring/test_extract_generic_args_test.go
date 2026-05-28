//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestExtractGenericArgs — 제네릭 타입에서 인자 추출 테스트
package spring

import "testing"

func TestExtractGenericArgs(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"PagedResponse<AlbumResponse>", "AlbumResponse"},
		{"AlbumResponse", ""},
		{"Map<String, Object>", "String, Object"},
	}
	for _, tt := range tests {
		got := extractGenericArgs(tt.input)
		if got != tt.want {
			t.Errorf("extractGenericArgs(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
