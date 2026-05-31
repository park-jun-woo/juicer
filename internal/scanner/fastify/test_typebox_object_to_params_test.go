//ff:func feature=scan type=test topic=fastify control=sequence
//ff:what typeBoxObjectToParams object → Param 슬라이스 변환 테스트
package fastify

import "testing"

func TestTypeBoxObjectToParams(t *testing.T) {
	fi := mustParse(t, []byte(`const q = { page: Type.Integer(), size: Type.Integer() };`))
	obj := findAllByType(fi.Root, "object")[0]
	params := typeBoxObjectToParams(obj, fi.Src)
	if len(params) != 2 || params[0].Name != "page" {
		t.Errorf("got %+v", params)
	}
	// non-object -> nil
	if typeBoxObjectToParams(nil, fi.Src) != nil {
		t.Error("nil should be nil")
	}
}
