//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestMethodPrefixedID 테스트
package scanner

import "testing"

func TestMethodPrefixedID(t *testing.T) {
	tests := []struct {
		method string
		id     string
		want   string
	}{
		{"GET", "login", "getLogin"},
		{"POST", "login", "postLogin"},
		{"DELETE", "user", "deleteUser"},
		{"", "login", "login"},
		{"GET", "", "get"},
	}
	for _, tt := range tests {
		t.Run(tt.method+"_"+tt.id, func(t *testing.T) {
			got := methodPrefixedID(tt.method, tt.id)
			if got != tt.want {
				t.Errorf("methodPrefixedID(%q,%q) = %q, want %q", tt.method, tt.id, got, tt.want)
			}
		})
	}
}
