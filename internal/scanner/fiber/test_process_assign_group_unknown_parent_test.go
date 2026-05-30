//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestProcessAssign_GroupUnknownParent 테스트
package fiber

import "testing"

func TestProcessAssign_GroupUnknownParent(t *testing.T) {
	src := `package m
func f() {
	sub := unknown.Group("/x")
}
`
	routers := map[string]*routerInfo{}
	for _, a := range assignStmts(t, src) {
		processAssign(a, "fiber", routers)
	}
	if _, ok := routers["sub"]; ok {
		t.Fatal("group with unknown parent should not register")
	}
}
