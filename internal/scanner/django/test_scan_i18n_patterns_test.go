//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what i18n_patterns(...) 래퍼로 감싼 urlpatterns 언래핑을 E2E로 검증한다
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_I18nPatterns(t *testing.T) {
	dir := t.TempDir()

	views := `from django.http import JsonResponse

def about(request):
    return JsonResponse({})
`
	os.WriteFile(filepath.Join(dir, "views.py"), []byte(views), 0o644)

	urls := `from django.urls import path
from django.conf.urls.i18n import i18n_patterns
from . import views

urlpatterns = i18n_patterns(
    path('about/', views.about),
)
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
	if !found["GET /about/"] {
		t.Errorf("missing GET /about/ from i18n_patterns wrapper")
	}
}
