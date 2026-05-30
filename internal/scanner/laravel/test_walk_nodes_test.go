//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestWalkNodes 테스트
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestWalkNodes(t *testing.T) {
	fi := mustParsePHP(t, `<?php foo(bar());`)
	count := 0
	walkNodes(fi.root, func(n *sitter.Node) {
		if n.Type() == "function_call_expression" {
			count++
		}
	})
	if count != 2 {
		t.Fatalf("expected 2 calls, got %d", count)
	}
}
