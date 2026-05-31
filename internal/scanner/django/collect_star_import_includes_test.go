//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what collectStarImportIncludes — 패키지 __init__.py 집계를 prefix 없는 include로 변환하는지 검증
package django

import "testing"

func TestCollectStarImportIncludes(t *testing.T) {
	src := []byte(`from .workspace import urlpatterns as workspace_urls
from .api import urlpatterns
from .helpers import something

urlpatterns = [*workspace_urls, *urlpatterns]
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	fi := fileInfo{relPath: "plane/app/urls/__init__.py", module: "plane.app.urls", src: src, root: root}

	entries := collectStarImportIncludes(fi)
	got := map[string]bool{}
	for _, e := range entries {
		if !e.isInclude {
			t.Errorf("entry not marked include: %+v", e)
		}
		got[e.includeModule] = true
	}
	if !got["plane.app.urls.workspace"] || !got["plane.app.urls.api"] {
		t.Errorf("missing aggregated submodules: %v", got)
	}
	if got["plane.app.urls.helpers"] {
		t.Error("helpers (not importing urlpatterns) should be ignored")
	}

	// Non-__init__ files produce nothing.
	notInit := fileInfo{relPath: "plane/app/urls/workspace.py", module: "plane.app.urls.workspace", src: src, root: root}
	if e := collectStarImportIncludes(notInit); e != nil {
		t.Errorf("non-__init__ file should yield no includes, got %+v", e)
	}
}
