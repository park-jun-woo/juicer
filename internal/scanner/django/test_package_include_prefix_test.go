//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what 패키지형 include의 prefix 합성과 서브모듈 루트 오인 방지를 검증
package django

import "testing"

func TestPackageIncludePrefix(t *testing.T) {
	rootUrls := mkFile(t, "proj/urls.py", "proj.urls", `
urlpatterns = [
    path("api/", include("app.urls")),
]
`)
	pkgInit := mkFile(t, "app/urls/__init__.py", "app.urls", `
from .workspace import urlpatterns as workspace_urls

urlpatterns = [*workspace_urls]
`)
	submodule := mkFile(t, "app/urls/workspace.py", "app.urls.workspace", `
urlpatterns = [
    path("workspaces/", view_fn),
]
`)
	files := []fileInfo{rootUrls, pkgInit, submodule}

	byModule := collectURLs(files)

	// The submodule must be marked as included (not a root).
	roots := findRootURLModules(byModule)
	for _, r := range roots {
		if r == "app.urls.workspace" {
			t.Errorf("submodule app.urls.workspace mistaken for a root: roots=%v", roots)
		}
	}

	// Full expansion yields the api/ prefix composed onto the submodule path.
	var paths []string
	for _, r := range roots {
		for _, e := range expandURLModule(r, "", byModule, map[string]bool{}) {
			paths = append(paths, e.pattern)
		}
	}
	found := false
	for _, p := range paths {
		if p == "/api/workspaces/" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected /api/workspaces/, got %v", paths)
	}
}
