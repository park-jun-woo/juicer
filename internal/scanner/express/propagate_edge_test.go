//ff:func feature=scan type=test control=sequence topic=express
//ff:what propagateEdge: 새 prefix추가 changed / 중복 no-change / 빈 pps
package express

import "testing"

func TestPropagateEdge_Changed(t *testing.T) {
	child := routerKey{file: "b", varName: "users"}
	prefixes := map[routerKey][]string{}
	e := routerEdge{child: child, seg: "/users"}
	if !propagateEdge(prefixes, e, []string{"/api"}) {
		t.Fatal("expected changed=true")
	}
	if len(prefixes[child]) != 1 || prefixes[child][0] != "/api/users" {
		t.Fatalf("got %v", prefixes[child])
	}
}

func TestPropagateEdge_NoChangeOnDuplicate(t *testing.T) {
	child := routerKey{file: "b", varName: "users"}
	prefixes := map[routerKey][]string{child: {"/api/users"}}
	e := routerEdge{child: child, seg: "/users"}
	if propagateEdge(prefixes, e, []string{"/api"}) {
		t.Fatal("expected changed=false for duplicate")
	}
}

func TestPropagateEdge_EmptyParents(t *testing.T) {
	child := routerKey{file: "b", varName: "users"}
	prefixes := map[routerKey][]string{}
	e := routerEdge{child: child, seg: "/users"}
	if propagateEdge(prefixes, e, nil) {
		t.Fatal("expected changed=false for empty parents")
	}
}
