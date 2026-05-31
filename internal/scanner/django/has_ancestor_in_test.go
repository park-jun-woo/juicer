//ff:func feature=scan type=test control=sequence topic=django
//ff:what hasAncestorIn — 전이적 상속 walk와 cycle 가드를 검증
package django

import "testing"

func TestHasAncestorIn(t *testing.T) {
	targets := map[string]bool{"ModelViewSet": true}

	// Direct parent in targets.
	if !hasAncestorIn([]string{"ModelViewSet"}, targets, nil) {
		t.Error("expected true for direct target parent")
	}

	// Transitive: V -> BaseViewSet -> ModelViewSet.
	idx := classIndex{
		"BaseViewSet": {"ModelViewSet"},
	}
	if !hasAncestorIn([]string{"BaseViewSet"}, targets, idx) {
		t.Error("expected true for transitive target ancestor")
	}

	// Unrelated chain stays false.
	if hasAncestorIn([]string{"object"}, targets, idx) {
		t.Error("expected false for unrelated parent")
	}

	// Cycle must not loop forever and stays false.
	cyclic := classIndex{
		"A": {"B"},
		"B": {"A"},
	}
	if hasAncestorIn([]string{"A"}, targets, cyclic) {
		t.Error("expected false for cyclic chain with no target")
	}
}
