//ff:func feature=scan type=test control=sequence topic=django
//ff:what buildClassIndex — 전 파일 클래스 이름→부모목록 인덱스 구축을 검증
package django

import "testing"

func TestBuildClassIndex(t *testing.T) {
	base := newTestFileInfo(t, `
class BaseViewSet(ModelViewSet):
    pass
`)
	view := newTestFileInfo(t, `
class ProjectViewSet(BaseViewSet):
    pass
`)
	idx := buildClassIndex([]fileInfo{base, view})

	if got := idx["BaseViewSet"]; len(got) != 1 || got[0] != "ModelViewSet" {
		t.Errorf("BaseViewSet parents = %v, want [ModelViewSet]", got)
	}
	if got := idx["ProjectViewSet"]; len(got) != 1 || got[0] != "BaseViewSet" {
		t.Errorf("ProjectViewSet parents = %v, want [BaseViewSet]", got)
	}
}
