//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseViewSetClass_Round5 테스트
package django

import "testing"

func TestParseViewSetClass_Round5(t *testing.T) {
	fi := newTestFileInfo(t, "class UserViewSet(ModelViewSet):\n    queryset = User.objects.all()\n")
	cls := djFirst(t, fi.root, "class_definition")
	vs := parseViewSetClass(cls, fi, nil)
	if vs == nil || vs.name != "UserViewSet" {
		t.Fatalf("viewset: %+v", vs)
	}
}
