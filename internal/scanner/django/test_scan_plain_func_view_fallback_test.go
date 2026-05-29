//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what @api_view 없는 순수 함수 뷰가 GET 폴백으로 잡히는지 E2E로 검증한다
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_PlainFuncViewFallback(t *testing.T) {
	dir := t.TempDir()

	views := `from django.http import JsonResponse

def home(request):
    return JsonResponse({})
`
	os.WriteFile(filepath.Join(dir, "views.py"), []byte(views), 0o644)

	urls := `from django.urls import path
from . import views
from django.contrib.auth import views as auth_views

urlpatterns = [
    path('', views.home),
    path('login/', auth_views.LoginView.as_view()),
]
`
	os.WriteFile(filepath.Join(dir, "urls.py"), []byte(urls), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
		t.Logf("found: %s %s (%s)", ep.Method, ep.Path, ep.Handler)
	}
	if !found["GET /"] {
		t.Errorf("missing GET / fallback for plain func view")
	}
	if !found["GET /login/"] {
		t.Errorf("missing GET /login/ fallback for CBV .as_view()")
	}
}
