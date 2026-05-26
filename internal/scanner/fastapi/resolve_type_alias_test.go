//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestResolveTypeAliases_Local 로컬 별칭 맵 구축 테스트
package fastapi

import "testing"

func TestResolveTypeAliases_Local(t *testing.T) {
	src := []byte(`
from typing import Annotated
from fastapi import Depends

SessionDep = Annotated[Session, Depends(get_db)]
CurrentUser = Annotated[User, Depends(get_current_user)]
TokenDep = Annotated[str, Depends(reusable_oauth2)]
NotAnAlias = SomeOtherType
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	files := []fileInfo{
		{absPath: "/app/deps.py", relPath: "deps.py", src: src, root: root},
	}
	m := resolveTypeAliases(files)

	tests := []struct {
		alias string
		want  string
	}{
		{"SessionDep", "get_db"},
		{"CurrentUser", "get_current_user"},
		{"TokenDep", "reusable_oauth2"},
	}
	for _, tc := range tests {
		if got := m[tc.alias]; got != tc.want {
			t.Errorf("alias %q: got %q, want %q", tc.alias, got, tc.want)
		}
	}
	if _, ok := m["NotAnAlias"]; ok {
		t.Error("NotAnAlias should not be in alias map")
	}
}
