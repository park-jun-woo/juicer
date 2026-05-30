//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyBodyParam 테스트
package quarkus

import "testing"

func TestClassifyBodyParam(t *testing.T) {
	ep := &endpointInfo{}
	classifyBodyParam("UserDto", "dto", ep)
	if ep.bodyType != "UserDto" || ep.bodyVarName != "dto" {
		t.Fatalf("got %+v", ep)
	}

	ep2 := &endpointInfo{}
	classifyBodyParam("String", "s", ep2)
	if ep2.bodyType != "" {
		t.Fatalf("primitive should be ignored: %+v", ep2)
	}

	ep3 := &endpointInfo{bodyType: "First"}
	classifyBodyParam("Second", "x", ep3)
	if ep3.bodyType != "First" {
		t.Fatalf("second body should be ignored: %+v", ep3)
	}
}
