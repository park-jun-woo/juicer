//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestIsAuthEndpoint 테스트
package scanner

import "testing"

func TestIsAuthEndpoint(t *testing.T) {
	tests := []struct {
		name string
		ep   Endpoint
		want bool
	}{
		{
			name: "authLevel auth_required",
			ep:   Endpoint{AuthLevel: "auth_required"},
			want: true,
		},
		{
			name: "authLevel public",
			ep:   Endpoint{AuthLevel: "public"},
			want: false,
		},
		{
			name: "middleware with auth keyword",
			ep:   Endpoint{Middleware: []string{"JwtAuthGuard"}},
			want: true,
		},
		{
			name: "middleware with current_user",
			ep:   Endpoint{Middleware: []string{"get_current_user"}},
			want: true,
		},
		{
			name: "middleware with verify_token",
			ep:   Endpoint{Middleware: []string{"verify_token"}},
			want: true,
		},
		{
			name: "middleware non-auth get_db",
			ep:   Endpoint{Middleware: []string{"get_db"}},
			want: false,
		},
		{
			name: "middleware non-auth SessionDep",
			ep:   Endpoint{Middleware: []string{"SessionDep"}},
			want: false,
		},
		{
			name: "no authLevel no middleware",
			ep:   Endpoint{},
			want: false,
		},
		{
			name: "authLevel public overrides middleware",
			ep:   Endpoint{AuthLevel: "public", Middleware: []string{"JwtAuthGuard"}},
			want: false,
		},
		{
			name: "mixed middleware auth and non-auth",
			ep:   Endpoint{Middleware: []string{"get_db", "verify_token"}},
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
