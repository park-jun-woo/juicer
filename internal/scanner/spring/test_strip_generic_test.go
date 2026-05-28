//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestStripGeneric — 제네릭 타입에서 래퍼 이름 분리 테스트
package spring

import "testing"

func TestStripGeneric(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"PagedResponse<AlbumResponse>", "PagedResponse"},
		{"AlbumResponse", "AlbumResponse"},
		{"Map<String, Object>", "Map"},
	}
	for _, tt := range tests {
		got := stripGeneric(tt.input)
		if got != tt.want {
			t.Errorf("stripGeneric(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
