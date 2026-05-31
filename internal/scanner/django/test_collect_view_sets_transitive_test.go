//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what collectViewSets — 커스텀 중간 베이스 경유 ViewSet 전이 인식을 검증
package django

import "testing"

func TestCollectViewSetsTransitive(t *testing.T) {
	base := newTestFileInfo(t, `
class BaseViewSet(ModelViewSet):
    pass
`)
	view := newTestFileInfo(t, `
class ProjectViewSet(BaseViewSet):
    serializer_class = ProjectSerializer
`)
	files := []fileInfo{base, view}
	idx := buildClassIndex(files)
	vs := collectViewSets(files, idx)

	var found bool
	for _, v := range vs {
		if v.name == "ProjectViewSet" {
			found = true
		}
	}
	if !found {
		t.Fatalf("ProjectViewSet not recognized transitively via BaseViewSet; got %+v", vs)
	}

	// Without the index the custom base is invisible (regression guard).
	if v := collectViewSets(files, nil); len(v) != 1 || v[0].name != "BaseViewSet" {
		t.Errorf("nil index should only see BaseViewSet, got %+v", v)
	}
}
