//ff:func feature=scan type=test topic=nestjs control=sequence
//ff:what resolveJobFile import 매핑 기반 타입 정의 파일 해석/부재 테스트
package nestjs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveJobFile(t *testing.T) {
	dir := t.TempDir()
	// referrer in dir, importing './user.dto'
	target := filepath.Join(dir, "user.dto.ts")
	if err := os.WriteFile(target, []byte("export class UserDto {}"), 0o644); err != nil {
		t.Fatal(err)
	}
	j := schemaJob{
		typeName:    "UserDto",
		imports:     map[string]string{"UserDto": "./user.dto"},
		referrer:    filepath.Join(dir, "controller.ts"),
		projectRoot: dir,
	}
	got := resolveJobFile(j)
	if got != target {
		t.Errorf("got %q, want %q", got, target)
	}
	// type not imported -> ""
	j2 := schemaJob{typeName: "X", imports: map[string]string{}, referrer: j.referrer, projectRoot: dir}
	if got := resolveJobFile(j2); got != "" {
		t.Errorf("unimported: got %q", got)
	}
}
