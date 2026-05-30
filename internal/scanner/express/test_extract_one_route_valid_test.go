//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneRoute_Valid 테스트
package express

import "testing"

func TestExtractOneRoute_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`r.get('/x', h);`))
	ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true})
	if ri == nil || ri.Method != "GET" || ri.Path != "/x" || ri.Router != "r" {
		t.Fatalf("got %+v", ri)
	}
}
