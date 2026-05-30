//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what stripLeadingPath 테스트
package fastapi

import "testing"

func TestStripLeadingPath(t *testing.T) {
	tests := []struct{ s, prefix, want string }{
		{"/api/v1/items", "/api", "/v1/items"}, // normal strip
		{"/api", "/api", ""},                   // exact match -> ""
		{"/api/v1", "", "/api/v1"},             // empty prefix -> unchanged
		{"/other/x", "/api", "/other/x"},       // prefix not found -> unchanged
		{"/apiextra", "/api", "/extra"},        // partial: trims then re-prefixes
	}
	for _, tt := range tests {
		if got := stripLeadingPath(tt.s, tt.prefix); got != tt.want {
			t.Errorf("stripLeadingPath(%q,%q) = %q, want %q", tt.s, tt.prefix, got, tt.want)
		}
	}
}
