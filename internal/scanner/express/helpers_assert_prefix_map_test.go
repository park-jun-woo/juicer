//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 테스트 헬퍼: resolveLocalRouterPrefixes 결과가 기대 prefix 맵과 일치하는지 검증한다
package express

import "testing"

func assertPrefixMap(t *testing.T, result []mountEntry, wantPath map[string]string) {
	t.Helper()
	got := make(map[string]string)
	for _, e := range result {
		got[e.filePath] = e.prefix
	}
	for fp, wantPrefix := range wantPath {
		if got[fp] != wantPrefix {
			t.Errorf("%s: want prefix %q, got %q", fp, wantPrefix, got[fp])
		}
	}
	for _, e := range result {
		if e.filePath == "" {
			t.Errorf("filePath='' entry should be removed, got varName=%s", e.varName)
		}
	}
}
