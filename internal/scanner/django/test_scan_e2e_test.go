//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what Django + DRF E2E 스캔 테스트 — ViewSet, APIView, @api_view 라우트를 통합 검증한다
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_E2E(t *testing.T) {
	dir := t.TempDir()

	// serializers.py
	serializersSrc := `from rest_framework import serializers

class UserSerializer(serializers.Serializer):
    name = serializers.CharField(max_length=100)
    email = serializers.EmailField()
    age = serializers.IntegerField(min_value=0, max_value=150)
`
	os.WriteFile(filepath.Join(dir, "serializers.py"), []byte(serializersSrc), 0o644)

	// views.py
	viewsSrc := `from rest_framework import viewsets, status
from rest_framework.decorators import api_view, action
from rest_framework.response import Response
from rest_framework.views import APIView

class UserViewSet(viewsets.ModelViewSet):
    serializer_class = UserSerializer

    @action(detail=True, methods=["post"])
    def activate(self, request, pk=None):
        return Response(status=status.HTTP_200_OK)

    @action(detail=False, methods=["get"])
    def recent(self, request):
        return Response([])

class HealthView(APIView):
    def get(self, request):
        return Response({"status": "ok"})

@api_view(["GET", "POST"])
def item_list(request):
    return Response([])
`
	os.WriteFile(filepath.Join(dir, "views.py"), []byte(viewsSrc), 0o644)

	// urls.py
	urlsSrc := `from django.urls import path
from rest_framework.routers import DefaultRouter
from . import views

router = DefaultRouter()
router.register(r"users", views.UserViewSet)

urlpatterns = [
    path("api/health/", views.HealthView.as_view()),
    path("api/items/", views.item_list),
]
`
	os.WriteFile(filepath.Join(dir, "urls.py"), []byte(urlsSrc), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	// Expected endpoints:
	// From router: list, create, retrieve, update, partial_update, destroy (6)
	// From router @action: activate (1), recent (1) = 2
	// From URLPatterns: HealthView.get (1), item_list GET (1), item_list POST (1)
	// Total: 11

	if len(result.Endpoints) < 8 {
		for i, ep := range result.Endpoints {
			t.Logf("  endpoint %d: %s %s (%s)", i, ep.Method, ep.Path, ep.Handler)
		}
		t.Fatalf("expected at least 8 endpoints, got %d", len(result.Endpoints))
	}

	// Verify a few key endpoints
	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		key := ep.Method + " " + ep.Path
		found[key] = true
		t.Logf("  found: %s %s (%s)", ep.Method, ep.Path, ep.Handler)

		verifyEndpointDetail(t, ep)
	}

	// Check specific expected endpoints
	expectedEndpoints := []string{
		"GET /users",
		"POST /users",
		"GET /users/{pk}",
	}
	for _, expected := range expectedEndpoints {
		if !found[expected] {
			t.Errorf("missing expected endpoint: %s", expected)
		}
	}
}
