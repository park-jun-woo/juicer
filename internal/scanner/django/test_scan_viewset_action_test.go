//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what ViewSet의 @action 데코레이터에서 커스텀 엔드포인트를 추출한다
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_ViewSetAction(t *testing.T) {
	dir := t.TempDir()

	viewsSrc := `from rest_framework import viewsets
from rest_framework.decorators import action
from rest_framework.response import Response

class UserViewSet(viewsets.ModelViewSet):
    serializer_class = UserSerializer

    @action(detail=True, methods=["post"])
    def activate(self, request, pk=None):
        return Response({"status": "activated"})

    @action(detail=False, methods=["get"])
    def recent(self, request):
        return Response([])
`
	os.WriteFile(filepath.Join(dir, "views.py"), []byte(viewsSrc), 0o644)

	urlsSrc := `from rest_framework.routers import DefaultRouter
router = DefaultRouter()
router.register(r"users", UserViewSet)
urlpatterns = router.urls
`
	os.WriteFile(filepath.Join(dir, "urls.py"), []byte(urlsSrc), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	// ModelViewSet: 6 default + 2 custom actions = 8
	if len(result.Endpoints) != 8 {
		for i, ep := range result.Endpoints {
			t.Logf("  endpoint %d: %s %s (%s)", i, ep.Method, ep.Path, ep.Handler)
		}
		t.Fatalf("expected 8 endpoints, got %d", len(result.Endpoints))
	}

	// Check for custom action endpoints
	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		key := ep.Method + " " + ep.Path
		found[key] = true
	}

	if !found["POST /users/{pk}/activate"] {
		t.Error("missing POST /users/{pk}/activate")
	}
	if !found["GET /users/recent"] {
		t.Error("missing GET /users/recent")
	}
}
