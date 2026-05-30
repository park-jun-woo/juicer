//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestBuildRequest_BodyAndQuery 테스트
package supafunc

import "testing"

func TestBuildRequest_BodyAndQuery(t *testing.T) {
	r := buildRequest([]string{"task", "status"}, []string{"limit"})
	if r == nil || r.Body == nil || len(r.Body.Fields) != 2 {
		t.Fatalf("body: %+v", r)
	}
	if len(r.Query) != 1 || r.Query[0].Name != "limit" {
		t.Fatalf("query: %+v", r.Query)
	}
}
