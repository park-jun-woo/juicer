//ff:func feature=scan type=test control=sequence dimension=1 topic=express
//ff:what 부모 prefix 없는 직접 마운트 테스트: prefix 변경 없음
package express

import "testing"

func TestBuildPrefixMapNoParent(t *testing.T) {
	mounts := []mountEntry{
		{prefix: "/users", varName: "usersRouter", filePath: "/app/routes/users.ts", sourceFile: "/app/app.ts"},
	}

	result := buildPrefixMap(mounts)

	got := result["/app/routes/users.ts"]
	if got != "/users" {
		t.Errorf("expected /users, got %s", got)
	}
}
