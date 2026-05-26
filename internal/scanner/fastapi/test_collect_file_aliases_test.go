//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestCollectFileAliases 단일 파일 별칭 수집 테스트
package fastapi

import "testing"

func TestCollectFileAliases(t *testing.T) {
	src := []byte(`
SessionDep = Annotated[Session, Depends(get_db)]
x = 42
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	result := collectFileAliases(root, src)
	if result["SessionDep"] != "get_db" {
		t.Errorf("expected get_db, got %q", result["SessionDep"])
	}
	if _, ok := result["x"]; ok {
		t.Error("x should not be in aliases")
	}
}
