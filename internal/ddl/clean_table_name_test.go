//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what cleanTableName 단위 테스트
package ddl

import "testing"

func TestCleanTableName(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"users", "users"},
		{"public.profiles", "profiles"},
		{`"UserEvents"`, "UserEvents"},
		{`public."UserEvents"`, "UserEvents"},
		{`"public"."UserEvents"`, "UserEvents"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := cleanTableName(tt.in)
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
