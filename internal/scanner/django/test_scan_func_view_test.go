//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what @api_view 함수 기반 뷰를 스캔한다
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_FuncView(t *testing.T) {
	dir := t.TempDir()

	viewsSrc := `from rest_framework.decorators import api_view
from rest_framework.response import Response

@api_view(["GET", "POST"])
def user_list(request):
    return Response([])

@api_view(["GET"])
def health_check(request):
    return Response({"status": "ok"})
`
	os.WriteFile(filepath.Join(dir, "views.py"), []byte(viewsSrc), 0o644)

	urlsSrc := `from django.urls import path
from . import views

urlpatterns = [
    path("api/users/", views.user_list),
    path("api/health/", views.health_check),
]
`
	os.WriteFile(filepath.Join(dir, "urls.py"), []byte(urlsSrc), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Endpoints) != 3 {
		for i, ep := range result.Endpoints {
			t.Logf("  endpoint %d: %s %s (%s)", i, ep.Method, ep.Path, ep.Handler)
		}
		t.Fatalf("expected 3 endpoints (2 for user_list, 1 for health_check), got %d", len(result.Endpoints))
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}

	for _, expected := range []string{"GET /api/users/", "POST /api/users/", "GET /api/health/"} {
		if !found[expected] {
			t.Errorf("missing expected endpoint: %s", expected)
		}
	}
}
