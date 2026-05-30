//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestResolveClassFields_Cached 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveClassFields_Cached(t *testing.T) {
	cache := map[string][]scanner.Field{"X": {{Name: "a", Type: "string"}}}
	got, err := resolveClassFields("/ignored", "X", "/root", cache)
	if err != nil || len(got) != 1 {
		t.Fatalf("got %+v err %v", got, err)
	}
}
