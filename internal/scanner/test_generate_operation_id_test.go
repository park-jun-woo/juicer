//ff:func feature=scan type=convert control=iteration dimension=1
//ff:what TestGenerateOperationID 테스트
package scanner

import (
	"testing"
)

func TestGenerateOperationID(t *testing.T) {
	tests := []struct {
		name string
		ep   Endpoint
		want string
	}{
		{
			name: "handler with file prefix",
			ep:   Endpoint{Handler: "handler.go:h.ListUsers", Method: "GET", Path: "/users"},
			want: "listUsers",
		},
		{
			name: "inline handler",
			ep:   Endpoint{Handler: "(inline)", Method: "GET", Path: "/api/v1/users"},
			want: "get_users",
		},
		{
			name: "empty handler",
			ep:   Endpoint{Handler: "", Method: "POST", Path: "/api/v1/items"},
			want: "post_items",
		},
		{
			name: "method with receiver",
			ep:   Endpoint{Handler: "h.CreateUser", Method: "POST", Path: "/users"},
			want: "createUser",
		},
		{
			name: "with trailing ()",
			ep:   Endpoint{Handler: "CreateUser()", Method: "POST", Path: "/users"},
			want: "createUser",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateOperationID(tt.ep)
			if got != tt.want {
				t.Errorf("generateOperationID() = %q, want %q", got, tt.want)
			}
		})
	}
}
