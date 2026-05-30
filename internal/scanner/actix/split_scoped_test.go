//ff:func feature=scan type=test control=sequence topic=actix
//ff:what splitScoped — :: 기준 분할 분기를 검증
package actix

import (
	"reflect"
	"testing"
)

func TestSplitScoped(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"a::b::c", []string{"a", "b", "c"}},
		{"single", []string{"single"}},
		{"web::scope", []string{"web", "scope"}},
		// trailing "::" leaves an empty trailing segment that is dropped.
		{"a::", []string{"a"}},
	}
	for _, c := range cases {
		got := splitScoped(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("splitScoped(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
