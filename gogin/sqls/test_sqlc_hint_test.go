//ff:func feature=ratchet type=session control=iteration dimension=1
//ff:what TestSqlcHint 테스트
package sqls

import (
	"testing"
)

func TestSqlcHint(t *testing.T) {
	tests := []struct {
		name string
		sk   *MethodSkeleton
		want string
	}{
		{
			name: "SELECT with slice return",
			sk:   &MethodSkeleton{CRUD: "SELECT", Returns: []string{"[]User"}},
			want: ":many",
		},
		{
			name: "SELECT without slice",
			sk:   &MethodSkeleton{CRUD: "SELECT", Returns: []string{"User"}},
			want: ":one",
		},
		{
			name: "INSERT with RETURNING",
			sk:   &MethodSkeleton{CRUD: "INSERT", SQLFragments: []string{"INSERT INTO users RETURNING id"}},
			want: ":one",
		},
		{
			name: "INSERT without RETURNING",
			sk:   &MethodSkeleton{CRUD: "INSERT"},
			want: ":exec",
		},
		{
			name: "UPDATE",
			sk:   &MethodSkeleton{CRUD: "UPDATE"},
			want: ":exec",
		},
		{
			name: "DELETE",
			sk:   &MethodSkeleton{CRUD: "DELETE"},
			want: ":exec",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sqlcHint(tt.sk)
			if got != tt.want {
				t.Errorf("sqlcHint() = %q, want %q", got, tt.want)
			}
		})
	}
}
