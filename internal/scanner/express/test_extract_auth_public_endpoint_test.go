//ff:func feature=scan type=test control=sequence topic=express
//ff:what 단위 테스트: 인증 미들웨어 없는 엔드포인트의 AuthLevel이 public인지 검증한다
package express

import "testing"

func TestExtractAuthMiddleware_PublicEndpoint(t *testing.T) {
	src := []byte(`
const router = express.Router();
router.get("/health", healthCheck);
`)
	fi := mustParse(t, src)
	routers := collectRouters(fi, nil)
	routes := extractRoutes(fi, routers)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	r := routes[0]
	if r.AuthLevel != "public" {
		t.Errorf("AuthLevel: want public, got %s", r.AuthLevel)
	}
	if len(r.Roles) != 0 {
		t.Errorf("Roles: want empty, got %v", r.Roles)
	}
}
