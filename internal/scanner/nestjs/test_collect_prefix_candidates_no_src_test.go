//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectPrefixCandidates_NoSrc 테스트
package nestjs

import "testing"

func TestCollectPrefixCandidates_NoSrc(t *testing.T) {
	dir := t.TempDir()
	cands := collectPrefixCandidates(dir)
	if len(cands) != 1 {
		t.Fatalf("expected just main path, got %v", cands)
	}
}
