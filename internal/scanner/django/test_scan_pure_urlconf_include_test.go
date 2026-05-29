//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what 순수 Django urlconf의 다단계 include + prefix 합성을 E2E로 검증한다
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_PureURLConfInclude(t *testing.T) {
	dir := t.TempDir()

	// root urls.py mounts blog app via include with a prefix
	rootSrc := `from django.urls import path, include

urlpatterns = [
    path('blog/', include('blog.urls')),
]
`
	os.WriteFile(filepath.Join(dir, "urls.py"), []byte(rootSrc), 0o644)

	// blog app
	os.MkdirAll(filepath.Join(dir, "blog"), 0o755)
	os.WriteFile(filepath.Join(dir, "blog", "__init__.py"), []byte(""), 0o644)
	blogViews := `from django.http import JsonResponse

def post_detail(request, pk):
    return JsonResponse({})
`
	os.WriteFile(filepath.Join(dir, "blog", "views.py"), []byte(blogViews), 0o644)
	blogUrls := `from django.urls import path
from . import views

urlpatterns = [
    path('posts/<int:pk>/', views.post_detail),
]
`
	os.WriteFile(filepath.Join(dir, "blog", "urls.py"), []byte(blogUrls), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
		t.Logf("found: %s %s (%s)", ep.Method, ep.Path, ep.Handler)
	}
	if !found["GET /blog/posts/{pk}/"] {
		t.Errorf("missing include-composed endpoint GET /blog/posts/{pk}/")
	}
}
