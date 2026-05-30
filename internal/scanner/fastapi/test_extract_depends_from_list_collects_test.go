//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDependsFromList_Collects 테스트
package fastapi

import "testing"

func TestExtractDependsFromList_Collects(t *testing.T) {
	list, src := firstList(t, []byte("x = [Depends(auth), Depends(log), other()]\n"))
	deps := extractDependsFromList(list, src)
	if len(deps) != 2 || deps[0] != "auth" || deps[1] != "log" {
		t.Fatalf("got %v", deps)
	}
}
