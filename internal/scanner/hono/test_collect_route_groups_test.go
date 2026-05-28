//ff:func feature=scan type=test control=sequence topic=hono
//ff:what app.route() 그룹 수집 테스트
package hono

import "testing"

func TestCollectRouteGroups(t *testing.T) {
	src := []byte(`
import { Hono } from "hono"
const app = new Hono()
const users = new Hono()
users.get("/", listUsers)
users.post("/", createUser)
app.route("/api/users", users)
`)
	fi := mustParse(t, src)
	vars := collectHonoVars(fi)
	groups := collectRouteGroups(fi, vars)
	if len(groups) != 1 {
		t.Fatalf("expected 1 group, got %d", len(groups))
	}
	if groups[0].Prefix != "/api/users" {
		t.Errorf("expected prefix /api/users, got %s", groups[0].Prefix)
	}
	if groups[0].SubAppName != "users" {
		t.Errorf("expected subAppName users, got %s", groups[0].SubAppName)
	}
}
