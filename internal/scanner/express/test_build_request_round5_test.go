//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestBuildRequest_Round5 테스트
package express

import "testing"

func TestBuildRequest_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express.Router();`))
	ctx := round5Ctx()

	r := routeInfo{Method: "GET", Path: "/users/:id", Handler: "h"}
	req := buildRequest(r, []string{"id"}, ctx, fi)
	if req == nil || len(req.PathParams) != 1 || req.PathParams[0].Name != "id" {
		t.Fatalf("request: %+v", req)
	}

	if got := buildRequest(routeInfo{Method: "GET", Path: "/x", Handler: "h"}, nil, ctx, fi); got != nil {
		t.Fatalf("expected nil request, got %+v", got)
	}
}
