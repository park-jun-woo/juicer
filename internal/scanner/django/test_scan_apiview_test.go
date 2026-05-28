//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what APIView 클래스에서 HTTP 메서드를 추출한다
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_APIView(t *testing.T) {
	dir := t.TempDir()

	viewsSrc := `from rest_framework.views import APIView
from rest_framework.response import Response

class UserDetailView(APIView):
    def get(self, request, pk):
        return Response({})

    def put(self, request, pk):
        return Response({})

    def delete(self, request, pk):
        return Response(status=204)
`
	os.WriteFile(filepath.Join(dir, "views.py"), []byte(viewsSrc), 0o644)

	urlsSrc := `from django.urls import path
from . import views

urlpatterns = [
    path("api/users/<int:pk>/", views.UserDetailView.as_view()),
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
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}

	methods := map[string]bool{}
	for _, ep := range result.Endpoints {
		methods[ep.Method] = true
		if ep.Path != "/api/users/{pk}/" {
			t.Errorf("expected /api/users/{pk}, got %s", ep.Path)
		}
		// Check path params
		if ep.Request == nil || len(ep.Request.PathParams) == 0 {
			t.Errorf("%s should have pk path param", ep.Method)
		} else if ep.Request.PathParams[0].Type != "integer" {
			t.Errorf("pk should be integer, got %s", ep.Request.PathParams[0].Type)
		}
	}

	for _, m := range []string{"GET", "PUT", "DELETE"} {
		if !methods[m] {
			t.Errorf("missing method %s", m)
		}
	}
}
