//ff:func feature=prisma type=test control=iteration dimension=1 topic=prisma
//ff:what mapDefaultExpr Prisma 기본식 → SQL 기본값 변환 테스트
package prisma

import "testing"

func TestMapDefaultExpr(t *testing.T) {
	cases := []struct{ in, want string }{
		{"autoincrement()", ""},
		{"cuid()", ""},
		{"cuid(2)", ""},
		{"now()", "now()"},
		{"true", "true"},
		{"false", "false"},
		{"uuid()", "gen_random_uuid()"},
		{"uuid(4)", "gen_random_uuid()"},
		{"uuid(7)", "gen_random_uuid()"},
		{`"hello"`, "'hello'"},
		{"42", "42"},
		{`"`, `"`},
	}
	for _, c := range cases {
		if got := mapDefaultExpr(c.in); got != c.want {
			t.Errorf("mapDefaultExpr(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
