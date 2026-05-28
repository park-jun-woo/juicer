//ff:func feature=scan type=test control=sequence topic=express
//ff:what tsconfig.json이 없는 프로젝트에서 에러 없이 빈 맵을 반환한다
package express

import "testing"

func TestLoadTsconfigPaths_None(t *testing.T) {
	dir := t.TempDir()
	aliases := loadTsconfigPaths(dir)
	if len(aliases) != 0 {
		t.Errorf("expected empty aliases, got %v", aliases)
	}
}
