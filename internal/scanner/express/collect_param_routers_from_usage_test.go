//ff:func feature=scan type=test control=sequence topic=express
//ff:what collectParamRoutersFromUsage: 각 continue 분기 + 수집 성공 검증
package express

import "testing"

func TestCollectParamRoutersFromUsage_Collects(t *testing.T) {
	src := []byte(`r.get('/x', h);`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromUsage(fi, routers)
	if !routers["r"] {
		t.Fatalf("expected 'r', got %v", routers)
	}
}

func TestCollectParamRoutersFromUsage_NonHttpMethod(t *testing.T) {
	// member call but method not in httpMethods
	src := []byte(`r.foo('/x');`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromUsage(fi, routers)
	if len(routers) != 0 {
		t.Fatalf("expected none for non-http method, got %v", routers)
	}
}

func TestCollectParamRoutersFromUsage_ExpressIgnored(t *testing.T) {
	src := []byte(`express.get('/x', h); module.get('/y'); exports.get('/z');`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromUsage(fi, routers)
	if len(routers) != 0 {
		t.Fatalf("expected express/module/exports ignored, got %v", routers)
	}
}

func TestCollectParamRoutersFromUsage_NonStringFirstArg(t *testing.T) {
	// first arg not a string -> not collected
	src := []byte(`r.use(mw);`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromUsage(fi, routers)
	if len(routers) != 0 {
		t.Fatalf("expected none for non-string first arg, got %v", routers)
	}
}

func TestCollectParamRoutersFromUsage_NoArgs(t *testing.T) {
	src := []byte(`r.get();`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromUsage(fi, routers)
	if len(routers) != 0 {
		t.Fatalf("expected none for empty args, got %v", routers)
	}
}

func TestCollectParamRoutersFromUsage_PlainCall(t *testing.T) {
	// call_expression without member_expression
	src := []byte(`get('/x', h);`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromUsage(fi, routers)
	if len(routers) != 0 {
		t.Fatalf("expected none for plain call, got %v", routers)
	}
}
