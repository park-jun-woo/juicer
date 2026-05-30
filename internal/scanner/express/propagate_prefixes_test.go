//ff:func feature=scan type=test control=sequence topic=express
//ff:what propagatePrefixes: 그래프 따라 prefix 수렴 전파
package express

import "testing"

func TestPropagatePrefixes_Converges(t *testing.T) {
	g := newTestGraph()
	app := routerKey{file: "a", varName: "app"}
	users := routerKey{file: "b", varName: "users"}
	admin := routerKey{file: "c", varName: "admin"}
	graphAddEdge(g, app, users, "/users")
	graphAddEdge(g, users, admin, "/admin")

	prefixes := map[routerKey][]string{app: {""}}
	propagatePrefixes(g, prefixes)

	if len(prefixes[users]) != 1 || prefixes[users][0] != "/users" {
		t.Fatalf("users prefix: %v", prefixes[users])
	}
	if len(prefixes[admin]) != 1 || prefixes[admin][0] != "/users/admin" {
		t.Fatalf("admin prefix: %v", prefixes[admin])
	}
}
