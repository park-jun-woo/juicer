//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestUnwrapReturnType_Round5 테스트
package dotnet

import "testing"

func TestUnwrapReturnType_Round5(t *testing.T) {
	cases := []struct {
		in    string
		inner string
		array bool
	}{
		{"ActionResult<Foo>", "Foo", false},
		{"ActionResult<List<Foo>>", "Foo", true},
		{"Task<Bar>", "Bar", false},
		{"List<X>", "X", true},
		{"IEnumerable<Y>", "Y", true},
		{"IList<Z>", "Z", true},
		{"ICollection<W>", "W", true},
		{"Thing[]", "Thing", true},
		{"Plain", "Plain", false},
	}
	for _, c := range cases {
		inner, arr := unwrapReturnType(c.in)
		if inner != c.inner || arr != c.array {
			t.Errorf("%q: got (%q,%v), want (%q,%v)", c.in, inner, arr, c.inner, c.array)
		}
	}
}
