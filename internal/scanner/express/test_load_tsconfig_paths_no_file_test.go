//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestLoadTsconfigPaths_NoFile 테스트
package express

import "testing"

func TestLoadTsconfigPaths_NoFile(t *testing.T) {
	dir := t.TempDir()
	if aliases := loadTsconfigPaths(dir); len(aliases) != 0 {
		t.Fatalf("expected empty when no tsconfig, got %v", aliases)
	}
}
