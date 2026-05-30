//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestBuildEndpointsFromRoute_And_OneEndpoint_Round5 테스트
package express

import "testing"

func TestBuildEndpointsFromRoute_And_OneEndpoint_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express.Router();`))
	ctx := round5Ctx()
	r := routeInfo{Method: "GET", Path: "/users/:id", Router: "r", Handler: "getUser", Line: 1}

	eps := buildEndpointsFromRoute(r, "/api", "routes.ts", ctx, fi)
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if eps[0].Method != "GET" || eps[0].Handler != "getUser" {
		t.Fatalf("endpoint: %+v", eps[0])
	}
	if eps[0].Path != "/api/users/{id}" {
		t.Errorf("path: %q", eps[0].Path)
	}

	ep := buildOneEndpoint("POST", "/api/users", r, "routes.ts", []string{}, ctx, fi)
	if ep.Method != "POST" {
		t.Fatalf("buildOneEndpoint method: %q", ep.Method)
	}

	rAll := routeInfo{Method: "all", Path: "/x", Handler: "h"}
	allEps := buildEndpointsFromRoute(rAll, "", "r.ts", ctx, fi)
	if len(allEps) != 5 {
		t.Fatalf("expected 5 endpoints for all, got %d", len(allEps))
	}
}
