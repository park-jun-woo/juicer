//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what TestUnquotePython_Round5 테스트
package django

import "testing"

func TestUnquotePython_Round5(t *testing.T) {
	cases := map[string]string{
		`"x"`:       "x",
		`'y'`:       "y",
		`f"z"`:      "z",
		`r'p'`:      "p",
		`"""abc"""`: "abc",
		`a`:         "a",
		``:          "",
	}
	for in, want := range cases {
		if got := unquotePython(in); got != want {
			t.Errorf("unquotePython(%q)=%q want %q", in, got, want)
		}
	}
}
