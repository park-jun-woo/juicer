//ff:func feature=scan type=test control=sequence topic=django
//ff:what include 대상이 아닌 루트 urlconf 모듈을 찾는지 검증한다
package django

import "testing"

func TestFindRootURLModules(t *testing.T) {
	byModule := map[string][]urlEntry{
		"config.urls": {{pattern: "", isInclude: true, includeModule: "blog.urls"}},
		"blog.urls":   {{pattern: "posts/", viewName: "PostView"}},
	}
	roots := findRootURLModules(byModule)
	if len(roots) != 1 || roots[0] != "config.urls" {
		t.Fatalf("expected [config.urls], got %v", roots)
	}
}
