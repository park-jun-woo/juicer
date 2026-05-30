//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestExtractGroupArgPrefix_Round5 테스트
package echo

import "testing"

func TestExtractGroupArgPrefix_Round5(t *testing.T) {

	ctx := &groupArgCtx{
		routers: map[string]*routerInfo{"authGroup": {prefix: "/auth"}},
	}
	arg := exprFrom(t, `authGroup`)
	prefix, ri, ok := extractGroupArgPrefix(arg, ctx)
	if !ok || prefix != "/auth" || ri == nil {
		t.Fatalf("var group: %q %v", prefix, ok)
	}

	if _, _, ok := extractGroupArgPrefix(exprFrom(t, `42`), ctx); ok {
		t.Fatal("literal should not be a group")
	}
}
