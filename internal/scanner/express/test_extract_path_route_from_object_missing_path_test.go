//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractPathRouteFromObject_MissingPath 테스트
package express

import "testing"

func TestExtractPathRouteFromObject_MissingPath(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { route: r };`))
	if e := extractPathRouteFromObject(firstObject(t, fi), fi.Src); e != nil {
		t.Fatalf("expected nil, got %+v", e)
	}
}
