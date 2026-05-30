//ff:func feature=scan type=test control=sequence topic=express
//ff:what collectParamRouters 분기 테스트: 타입 어노테이션 경로와 사용 폴백 경로
package express

import "testing"

func TestCollectParamRouters_FromType(t *testing.T) {
	src := []byte(`function setup(r: Router) { r.get('/x', h); }`)
	fi := mustParse(t, src)
	routers := collectParamRouters(fi)
	if !routers["r"] {
		t.Fatalf("expected router 'r' from type annotation, got %v", routers)
	}
}

func TestCollectParamRouters_FromUsageFallback(t *testing.T) {
	// no Router type annotation -> from_type yields nothing -> usage fallback runs
	src := []byte(`function setup(r) { r.get('/x', h); }`)
	fi := mustParse(t, src)
	routers := collectParamRouters(fi)
	if !routers["r"] {
		t.Fatalf("expected router 'r' from usage fallback, got %v", routers)
	}
}

func TestCollectParamRouters_Empty(t *testing.T) {
	src := []byte(`function setup() { return 1; }`)
	fi := mustParse(t, src)
	routers := collectParamRouters(fi)
	if len(routers) != 0 {
		t.Fatalf("expected no routers, got %v", routers)
	}
}
