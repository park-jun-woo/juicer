//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what joinParts 테스트
package nestjs

import "testing"

func TestJoinParts_Cases(t *testing.T) {
	cases := []struct {
		parts []string
		want  string
	}{
		{[]string{}, "/"},
		{[]string{"users"}, "users"},
		{[]string{"users", ":id"}, "users/:id"},
		{[]string{"", "users", ""}, "users"},
		{[]string{"api//v1"}, "api/v1"},
	}
	for _, c := range cases {
		got := joinParts(c.parts...)
		if got != c.want {
			t.Errorf("joinParts(%v) = %q, want %q", c.parts, got, c.want)
		}
	}
}
