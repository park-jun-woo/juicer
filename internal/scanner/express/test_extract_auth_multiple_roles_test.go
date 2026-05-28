//ff:func feature=scan type=test control=sequence topic=express
//ff:what 단위 테스트: authorize('admin', 'manager') 미들웨어에서 복수 역할을 추출하는지 검증한다
package express

import "testing"

func TestExtractAuthMiddleware_AuthWithMultipleRoles(t *testing.T) {
	src := []byte(`
const router = express.Router();
router.put("/settings", authenticate, authorize('admin', 'manager'), updateSettings);
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
	if len(r.Roles) != 2 {
		t.Fatalf("Roles: want 2, got %d: %v", len(r.Roles), r.Roles)
	}
	if r.Roles[0] != "admin" {
		t.Errorf("Roles[0]: want admin, got %s", r.Roles[0])
	}
	if r.Roles[1] != "manager" {
		t.Errorf("Roles[1]: want manager, got %s", r.Roles[1])
	}
}
