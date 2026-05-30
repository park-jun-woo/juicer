//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what TestExtractHandlerName_Round5 테스트
package express

import "testing"

func TestExtractHandlerName_Round5(t *testing.T) {
	cases := map[string]string{
		`handler`:    "handler",
		`obj.method`: "obj.method",
		`() => {}`:   "(anonymous)",
		`wrap(h)`:    "wrap",
	}
	for src, want := range cases {
		f := mustParse(t, []byte("x = "+src+";"))
		node := rhsExpr(t, f)
		if got := extractHandlerName(node, f.Src); got != want {
			t.Errorf("extractHandlerName(%q)=%q want %q", src, got, want)
		}
	}
}
