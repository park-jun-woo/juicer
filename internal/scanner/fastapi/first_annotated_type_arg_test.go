//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what firstAnnotatedTypeArg: comma 구분 / 단일 닫힘 / 비Annotated / 미닫힘
package fastapi

import "testing"

func TestFirstAnnotatedTypeArg(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Annotated[User, Depends(get_user)]", "User"},        // comma at depth 0
		{"Annotated[Dict[str, int], Body()]", "Dict[str, int]"}, // nested brackets
		{"Annotated[User]", "User"},                            // closing bracket -> depth<0
		{"int", ""},                                            // not Annotated
		{"Annotated[unterminated", ""},                         // no close, no comma
	}
	for _, c := range cases {
		if got := firstAnnotatedTypeArg(c.in); got != c.want {
			t.Errorf("firstAnnotatedTypeArg(%q)=%q want %q", c.in, got, c.want)
		}
	}
}
