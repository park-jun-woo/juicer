//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestPrefixedOperationID 테스트
package scanner

import "testing"

func TestPrefixedOperationID(t *testing.T) {
	tests := []struct {
		name string
		ep   Endpoint
		id   string
		want string
	}{
		{
			name: "path segment differs from id uses path prefix",
			ep:   Endpoint{Method: "GET", Path: "/api/v1/categories"},
			id:   "findAll",
			want: "categoriesFindAll",
		},
		{
			name: "path segment equals id leading token uses method prefix",
			ep:   Endpoint{Method: "POST", Path: "/login"},
			id:   "login",
			want: "postLogin",
		},
		{
			name: "empty path segment uses method prefix",
			ep:   Endpoint{Method: "GET", Path: "/{id}"},
			id:   "login",
			want: "getLogin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prefixedOperationID(tt.ep, tt.id)
			if got != tt.want {
				t.Errorf("prefixedOperationID() = %q, want %q", got, tt.want)
			}
		})
	}
}
