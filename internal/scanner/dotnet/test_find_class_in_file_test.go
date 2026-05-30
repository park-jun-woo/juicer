//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestFindClassInFile 테스트
package dotnet

import "testing"

func TestFindClassInFile(t *testing.T) {
	fi := csFileInfo(t, `class UserDto { public string Name { get; set; } }`)
	fields := findClassInFile(fi, "UserDto")
	if len(fields) != 1 {
		t.Fatalf("got %+v", fields)
	}
	if findClassInFile(fi, "Missing") != nil {
		t.Fatal("missing should be nil")
	}
}
