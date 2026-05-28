//ff:func feature=scan type=test control=sequence topic=express
//ff:what 단위 테스트: 권한 미들웨어만 있어도 AuthLevel이 auth_required인지 검증한다
package express

import "testing"

func TestExtractAuthMiddleware_RoleOnlyImpliesAuth(t *testing.T) {
	src := []byte(`
const router = express.Router();
router.post("/admin/reset", authorize('admin'), resetSystem);
`)
	fi := mustParse(t, src)
	routers := collectRouters(fi, nil)
	routes := extractRoutes(fi, routers)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	r := routes[0]
	if r.AuthLevel != "auth_required" {
		t.Errorf("AuthLevel: want auth_required, got %s", r.AuthLevel)
	}
	if len(r.Roles) != 1 || r.Roles[0] != "admin" {
		t.Errorf("Roles: want [admin], got %v", r.Roles)
	}
}
