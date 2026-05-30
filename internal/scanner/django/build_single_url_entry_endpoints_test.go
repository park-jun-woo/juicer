//ff:func feature=scan type=test control=selection topic=django
//ff:what buildSingleURLEntryEndpoints — View 유형별 엔드포인트 생성 분기를 검증
package django

import "testing"

func TestBuildSingleURLEntry_EmptyView(t *testing.T) {
	eps := buildSingleURLEntryEndpoints(urlEntry{viewName: ""}, nil, nil, nil, nil)
	if eps != nil {
		t.Fatalf("expected nil for empty view name, got %+v", eps)
	}
}

func TestBuildSingleURLEntry_ViewSet(t *testing.T) {
	entry := urlEntry{pattern: "users/", viewName: "UserViewSet"}
	vs := []viewsetInfo{{name: "UserViewSet", parents: []string{"ModelViewSet"}, file: "v.py"}}
	eps := buildSingleURLEntryEndpoints(entry, vs, nil, nil, map[string]serializerInfo{})
	if len(eps) == 0 {
		t.Fatal("expected viewset endpoints")
	}
}

func TestBuildSingleURLEntry_APIView(t *testing.T) {
	entry := urlEntry{pattern: "ping/", viewName: "PingView"}
	av := []apiviewInfo{{name: "PingView", methods: []string{"GET"}, file: "v.py"}}
	eps := buildSingleURLEntryEndpoints(entry, nil, av, nil, map[string]serializerInfo{})
	if len(eps) == 0 {
		t.Fatal("expected apiview endpoints")
	}
}

func TestBuildSingleURLEntry_FuncView(t *testing.T) {
	entry := urlEntry{pattern: "health/", viewName: "health"}
	fv := []funcViewInfo{{name: "health", methods: []string{"GET"}, file: "v.py"}}
	eps := buildSingleURLEntryEndpoints(entry, nil, nil, fv, map[string]serializerInfo{})
	if len(eps) == 0 {
		t.Fatal("expected funcview endpoints")
	}
}

func TestBuildSingleURLEntry_Plain(t *testing.T) {
	// No matching view -> plain fallback.
	entry := urlEntry{pattern: "misc/", viewName: "UnknownView"}
	eps := buildSingleURLEntryEndpoints(entry, nil, nil, nil, map[string]serializerInfo{})
	if len(eps) == 0 {
		t.Fatal("expected plain endpoints fallback")
	}
}
