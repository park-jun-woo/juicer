//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestRunPrefixPass_NoChange 테스트
package hono

import "testing"

func TestRunPrefixPass_NoChange(t *testing.T) {

	prefixMap := map[string]string{}
	if runPrefixPass(nil, prefixMap, map[string]map[string]bool{}, map[string]map[string]string{}) {
		t.Fatal("expected no change for empty groups")
	}
}
