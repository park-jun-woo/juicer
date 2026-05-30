//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractMiddlewareNameForAuth_Round5 테스트
package express

import "testing"

func TestExtractMiddlewareNameForAuth_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`r.get('/x', requireAuth, handler);`))
	id := exFirst(t, fi, "identifier")
	name, ok := extractMiddlewareNameForAuth(id, fi.Src)
	_ = name
	_ = ok
}
