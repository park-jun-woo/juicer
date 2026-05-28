//ff:func feature=scan type=test control=sequence topic=express
//ff:what 단위 테스트: authenticate 미들웨어가 있으면 AuthLevel이 auth_required인지 검증한다
package express

import "testing"

func TestExtractAuthMiddleware_AuthRequired(t *testing.T) {
	src := []byte(`
const router = express.Router();
router.get("/profile", authenticate, getProfile);
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
	if len(r.Roles) != 0 {
		t.Errorf("Roles: want empty, got %v", r.Roles)
	}
}
