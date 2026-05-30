//ff:func feature=scan type=test control=sequence
//ff:what TestWellKnownType_TimeTime 테스트
package fiber

import "testing"

func TestWellKnownType_TimeTime(t *testing.T) {
	src := `package m
import "time"
var T time.Time
`
	named := namedTypeOf(t, src, "T")
	if name, ok := wellKnownType(named); !ok || name != "time.Time" {
		t.Fatalf("time.Time: %q %v", name, ok)
	}
}
