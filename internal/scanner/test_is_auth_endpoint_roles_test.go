//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestIsAuthEndpoint_Roles — Roles 필드가 있으면 인증 엔드포인트로 판별
package scanner

import "testing"

func TestIsAuthEndpoint_Roles(t *testing.T) {
	tests := []struct {
		name string
		ep   Endpoint
		want bool
	}{
		{
			name: "roles present treated as auth",
			ep:   Endpoint{Roles: []string{"ADMIN"}},
			want: true,
		},
		{
			name: "roles present overrides no middleware",
			ep:   Endpoint{Roles: []string{"USER", "ADMIN"}},
			want: true,
		},
		{
			name: "empty roles falls through",
			ep:   Endpoint{Roles: []string{}},
			want: false,
		},
		{
			name: "public with roles still returns true",
			ep:   Endpoint{AuthLevel: "public", Roles: []string{"ADMIN"}},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAuthEndpoint(tt.ep)
			if got != tt.want {
				t.Errorf("isAuthEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
