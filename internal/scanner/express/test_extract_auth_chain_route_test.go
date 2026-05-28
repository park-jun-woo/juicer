//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 단위 테스트: 체인 라우트에서 메서드별 인증/역할을 추출하는지 검증한다
package express

import "testing"

func TestExtractAuthMiddleware_ChainRoute(t *testing.T) {
	src := []byte(`
const router = express.Router();
router.route("/:id").get(authenticate, getUser).delete(authenticate, authorize('admin'), deleteUser);
`)
	fi := mustParse(t, src)
	routers := collectRouters(fi, nil)
	routes := extractRoutes(fi, routers)
	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(routes))
	}

	found := map[string]routeInfo{}
	for _, r := range routes {
		found[r.Method] = r
	}

	get := found["GET"]
	if get.AuthLevel != "auth_required" {
		t.Errorf("GET AuthLevel: want auth_required, got %s", get.AuthLevel)
	}
	if len(get.Roles) != 0 {
		t.Errorf("GET Roles: want empty, got %v", get.Roles)
	}

	del := found["DELETE"]
	if del.AuthLevel != "auth_required" {
		t.Errorf("DELETE AuthLevel: want auth_required, got %s", del.AuthLevel)
	}
	if len(del.Roles) != 1 || del.Roles[0] != "admin" {
		t.Errorf("DELETE Roles: want [admin], got %v", del.Roles)
	}
}
