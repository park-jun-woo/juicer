//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what tryParseIncludeRouter 테스트
package fastapi

import "testing"

func TestTryParseIncludeRouter(t *testing.T) {
	src := []byte("app.include_router(router, prefix='/v1')\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	calls := findAllByType(root, "call")
	found := false
	for _, call := range calls {
		inc := tryParseIncludeRouter(call, src)
		if inc == nil {
			continue
		}
		found = true
		if inc.parentVar != "app" || inc.childVar != "router" || inc.extraPrefix != "/v1" {
			t.Fatalf("got %+v", inc)
		}
	}
	if !found {
		t.Fatal("did not find include_router call")
	}
}
